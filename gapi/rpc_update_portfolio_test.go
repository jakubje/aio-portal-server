package gapi

import (
	"context"
	"database/sql"
	"github.com/jakub/aioportal/server/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	mockdb "github.com/jakub/aioportal/server/db/mock"
	db "github.com/jakub/aioportal/server/db/sqlc"
	"github.com/jakub/aioportal/server/pb"
	"github.com/jakub/aioportal/server/token"
	"github.com/stretchr/testify/require"
)

func TestUpdatePortfolioAPI(t *testing.T) {
	user, _ := randomUserPassword(t)
	portfolio := randomPortfolio(t, user)
	newName := util.RandomString(7)

	testCases := []struct {
		name          string
		req           *pb.UpdatePortfolioRequest
		buildStubs    func(store *mockdb.MockStore)
		buildContext  func(t *testing.T, tokenMaker token.Maker) context.Context
		checkResponse func(t *testing.T, res *pb.UpdatePortfolioResponse, err error)
	}{{
		name: "OK",
		req: &pb.UpdatePortfolioRequest{
			Id:   portfolio.ID,
			Name: newName,
		},
		buildStubs: func(store *mockdb.MockStore) {
			arg := db.UpdatePortfolioParams{
				ID:        portfolio.ID,
				AccountID: user.ID,
				Name:      newName,
			}

			updatedPortfolio := db.Portfolio{
				ID:         portfolio.ID,
				AccountID:  user.ID,
				Name:       newName,
				Holdings:   0,
				Change24h:  0,
				ProfitLoss: 0,
			}

			store.EXPECT().
				UpdatePortfolio(gomock.Any(), gomock.Eq(arg)).
				Times(1).
				Return(updatedPortfolio, nil)

		},
		buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
			return newContextWithBearerToken(t, tokenMaker, user.ID, user.Email, time.Minute)

		},
		checkResponse: func(t *testing.T, res *pb.UpdatePortfolioResponse, err error) {
			require.NoError(t, err)
			require.NotNil(t, res)
			updatedPortfolio := res.GetPortfolio()
			require.Equal(t, newName, updatedPortfolio.Name)
			require.Equal(t, portfolio.AccountID, updatedPortfolio.AccountId)
		},
	},
		{
			name: "InternalError",
			req: &pb.UpdatePortfolioRequest{
				Id:   portfolio.ID,
				Name: newName,
			},
			buildStubs: func(store *mockdb.MockStore) {

				store.EXPECT().
					UpdatePortfolio(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Portfolio{}, sql.ErrConnDone)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, user.ID, user.Email, time.Minute)

			},
			checkResponse: func(t *testing.T, res *pb.UpdatePortfolioResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.Internal, st.Code())
			},
		},
		{
			name: "InvalidName",
			req: &pb.UpdatePortfolioRequest{
				Id:   portfolio.ID,
				Name: "",
			},
			buildStubs: func(store *mockdb.MockStore) {

				store.EXPECT().
					UpdatePortfolio(gomock.Any(), gomock.Any()).
					Times(0)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, user.ID, user.Email, time.Minute)

			},
			checkResponse: func(t *testing.T, res *pb.UpdatePortfolioResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.InvalidArgument, st.Code())
			},
		},
		{
			name: "ExpiredToken",
			req: &pb.UpdatePortfolioRequest{
				Id:   portfolio.ID,
				Name: newName,
			},
			buildStubs: func(store *mockdb.MockStore) {

				store.EXPECT().
					UpdatePortfolio(gomock.Any(), gomock.Any()).
					Times(0)

			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, user.ID, user.Email, -time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.UpdatePortfolioResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.Unauthenticated, st.Code())
			},
		},
		{
			name: "NoAuthorization",
			req: &pb.UpdatePortfolioRequest{
				Id:   portfolio.ID,
				Name: portfolio.Name,
			},
			buildStubs: func(store *mockdb.MockStore) {

				store.EXPECT().
					UpdatePortfolio(gomock.Any(), gomock.Any()).
					Times(0)

			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return context.Background()
			},
			checkResponse: func(t *testing.T, res *pb.UpdatePortfolioResponse, err error) {
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
			res, err := server.UpdatePortfolio(ctx, tc.req)
			tc.checkResponse(t, res, err)
		})

	}
}
