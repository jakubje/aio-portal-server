package gapi

import (
	"context"
	"database/sql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	mockdb "github.com/jakub/aioportal/server/db/mock"
	db "github.com/jakub/aioportal/server/db/sqlc"
	"github.com/jakub/aioportal/server/pb"
	"github.com/jakub/aioportal/server/token"
	"github.com/jakub/aioportal/server/util"
	"github.com/stretchr/testify/require"
)

func randomPortfolio(t *testing.T, user db.User) db.Portfolio {

	portfolio := db.Portfolio{
		AccountID:  user.ID,
		Name:       util.RandomString(6),
		Holdings:   0,
		Change24h:  0,
		ProfitLoss: 0,
	}

	return portfolio
}

func TestCreatePortfolioAPI(t *testing.T) {
	user, _ := randomUserPassword(t)
	portfolio := randomPortfolio(t, user)

	testCases := []struct {
		name          string
		req           *pb.CreatePortfolioRequest
		buildStubs    func(store *mockdb.MockStore)
		buildContext  func(t *testing.T, tokenMaker token.Maker) context.Context
		checkResponse func(t *testing.T, res *pb.CreatePortfolioResponse, err error)
	}{{
		name: "OK",
		req: &pb.CreatePortfolioRequest{
			Name: portfolio.Name,
		},
		buildStubs: func(store *mockdb.MockStore) {
			arg := db.CreatePortfolioParams{
				Name: portfolio.Name,
			}

			store.EXPECT().
				CreatePortfolio(gomock.Any(), gomock.Eq(arg)).
				Times(1).
				Return(portfolio, nil)

		},
		buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
			return newContextWithBearerToken(t, tokenMaker, user.ID, user.Email, time.Minute)

		},
		checkResponse: func(t *testing.T, res *pb.CreatePortfolioResponse, err error) {
			require.NoError(t, err)
			require.NotNil(t, res)
			createdPortfolio := res.GetPortfolio()
			require.Equal(t, portfolio.Name, createdPortfolio.Name)
			require.Equal(t, portfolio.AccountID, createdPortfolio.AccountId)
		},
	},
		{
			name: "InternalError",
			req: &pb.CreatePortfolioRequest{
				Name: portfolio.Name,
			},
			buildStubs: func(store *mockdb.MockStore) {

				store.EXPECT().
					CreatePortfolio(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Portfolio{}, sql.ErrConnDone)

			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, user.ID, user.Email, time.Minute)

			},
			checkResponse: func(t *testing.T, res *pb.CreatePortfolioResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.Internal, st.Code())
			},
		},
		{
			name: "InvalidName",
			req: &pb.CreatePortfolioRequest{
				Name: "",
			},
			buildStubs: func(store *mockdb.MockStore) {

				store.EXPECT().
					CreatePortfolio(gomock.Any(), gomock.Any()).
					Times(0)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, user.ID, user.Email, time.Minute)

			},
			checkResponse: func(t *testing.T, res *pb.CreatePortfolioResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.InvalidArgument, st.Code())
			},
		},
		{
			name: "ExpiredToken",
			req: &pb.CreatePortfolioRequest{
				Name: portfolio.Name,
			},
			buildStubs: func(store *mockdb.MockStore) {

				store.EXPECT().
					CreatePortfolio(gomock.Any(), gomock.Any()).
					Times(0)

			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, user.ID, user.Email, -time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.CreatePortfolioResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.Unauthenticated, st.Code())
			},
		},
		{
			name: "NoAuthorization",
			req: &pb.CreatePortfolioRequest{
				Name: portfolio.Name,
			},
			buildStubs: func(store *mockdb.MockStore) {

				store.EXPECT().
					CreatePortfolio(gomock.Any(), gomock.Any()).
					Times(0)

			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return context.Background()
			},
			checkResponse: func(t *testing.T, res *pb.CreatePortfolioResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.Unauthenticated, st.Code())
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			storeCtrl := gomock.NewController(t)
			defer storeCtrl.Finish()

			store := mockdb.NewMockStore(storeCtrl)

			tc.buildStubs(store)
			server := newTestServer(t, store, nil)

			ctx := tc.buildContext(t, server.tokenMaker)
			res, err := server.CreatePortfolio(ctx, tc.req)
			tc.checkResponse(t, res, err)
		})

	}
}
