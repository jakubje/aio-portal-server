package gapi

import (
	"context"
	"github.com/jackc/pgx/v5"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/jackc/pgx/v5/pgtype"
	mockdb "github.com/jakub/aioportal/server/db/mock"
	db "github.com/jakub/aioportal/server/db/sqlc"
	"github.com/jakub/aioportal/server/pb"
	"github.com/jakub/aioportal/server/token"
	"github.com/jakub/aioportal/server/util"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestUpdateUserAPI(t *testing.T) {
	user, _ := randomUserPassword(t)

	newName := util.RandomString(6)
	newEmail := util.RandomEmail()
	invalidEmail := "invalid-email"

	testCases := []struct {
		name          string
		req           *pb.UpdateUserRequest
		buildStubs    func(store *mockdb.MockStore)
		buildContext  func(t *testing.T, tokenMaker token.Maker) context.Context
		checkResponse func(t *testing.T, res *pb.UpdateUserResponse, err error)
	}{{
		name: "OK",
		req: &pb.UpdateUserRequest{
			Id:    user.ID,
			Email: &newEmail,
			Name:  &newName,
		},
		buildStubs: func(store *mockdb.MockStore) {
			arg := db.UpdateUserParams{
				ID: user.ID,
				Name: pgtype.Text{
					String: newName,
					Valid:  true,
				},
				Email: pgtype.Text{
					String: newEmail,
					Valid:  true,
				},
			}
			updatedUser := db.User{
				ID:                user.ID,
				Email:             newEmail,
				Name:              newName,
				LastName:          user.LastName,
				Password:          user.Password,
				PasswordChangedAt: user.PasswordChangedAt,
				CreatedAt:         user.CreatedAt,
				IsEmailVerified:   user.IsEmailVerified,
			}
			store.EXPECT().
				UpdateUser(gomock.Any(), gomock.Eq(arg)).
				Times(1).
				Return(updatedUser, nil)

		},
		buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
			return newContextWithBearerToken(t, tokenMaker, user.ID, newEmail, time.Minute)

		},
		checkResponse: func(t *testing.T, res *pb.UpdateUserResponse, err error) {
			require.NoError(t, err)
			require.NotNil(t, res)
			updatedUser := res.GetUser()
			require.Equal(t, user.ID, updatedUser.ID)
			require.Equal(t, newName, updatedUser.Name)
			require.Equal(t, newEmail, updatedUser.Email)
		},
	},
		{
			name: "UserNotFound",
			req: &pb.UpdateUserRequest{
				Id:    user.ID,
				Email: &newEmail,
				Name:  &newName,
			},
			buildStubs: func(store *mockdb.MockStore) {

				store.EXPECT().
					UpdateUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.User{}, pgx.ErrNoRows)

			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, user.ID, newEmail, time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.UpdateUserResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.NotFound, st.Code())
			},
		},
		{
			name: "ExpiredToken",
			req: &pb.UpdateUserRequest{
				Id:    user.ID,
				Email: &newEmail,
				Name:  &newName,
			},
			buildStubs: func(store *mockdb.MockStore) {

				store.EXPECT().
					UpdateUser(gomock.Any(), gomock.Any()).
					Times(0)

			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, user.ID, newEmail, -time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.UpdateUserResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.Unauthenticated, st.Code())
			},
		},
		{
			name: "NoAuthorization",
			req: &pb.UpdateUserRequest{
				Id:    user.ID,
				Email: &newEmail,
				Name:  &newName,
			},
			buildStubs: func(store *mockdb.MockStore) {

				store.EXPECT().
					UpdateUser(gomock.Any(), gomock.Any()).
					Times(0)

			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return context.Background()
			},
			checkResponse: func(t *testing.T, res *pb.UpdateUserResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.Unauthenticated, st.Code())
			},
		},
		{
			name: "InvalidEmail",
			req: &pb.UpdateUserRequest{
				Id:    user.ID,
				Email: &invalidEmail,
				Name:  &newName,
			},
			buildStubs: func(store *mockdb.MockStore) {

				store.EXPECT().
					UpdateUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, user.ID, invalidEmail, time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.UpdateUserResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.InvalidArgument, st.Code())
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
			res, err := server.UpdateUser(ctx, tc.req)
			tc.checkResponse(t, res, err)
		})

	}
}
