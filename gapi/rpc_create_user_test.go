package gapi

import (
	"context"
	"github.com/jackc/pgx/v5"

	"fmt"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	mockdb "github.com/jakub/aioportal/server/db/mock"
	db "github.com/jakub/aioportal/server/db/sqlc"
	"github.com/jakub/aioportal/server/pb"
	"github.com/jakub/aioportal/server/util"
	"github.com/jakub/aioportal/server/worker"
	mockwk "github.com/jakub/aioportal/server/worker/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type eqCreateUserTxParamsMatcher struct {
	arg      db.CreateUserTxParams
	password string
	user     db.User
}

func (expected eqCreateUserTxParamsMatcher) Matches(x interface{}) bool {
	actualArg, ok := x.(db.CreateUserTxParams)
	if !ok {
		return false
	}

	err := util.CheckPassword(expected.password, actualArg.Password)
	if err != nil {
		return false
	}

	expected.arg.Password = actualArg.Password
	if !reflect.DeepEqual(expected.arg.CreateUserParams, actualArg.CreateUserParams) {
		return false
	}
	err = actualArg.AfterCreate(expected.user)
	return err == nil
}

func (e eqCreateUserTxParamsMatcher) String() string {
	return fmt.Sprintf("matches arg %v and password %v", e.arg, e.password)
}

func EqCreateUserTxParams(arg db.CreateUserTxParams, password string, user db.User) gomock.Matcher {
	return eqCreateUserTxParamsMatcher{arg, password, user}
}

func randomUserPassword(t *testing.T) (user db.User, password string) {
	password = util.RandomString(8)
	hashedPassword, err := util.HashPasswod(password)
	require.NoError(t, err)

	user = db.User{
		Email:    util.RandomEmail(),
		Name:     util.RandomString(6),
		LastName: util.RandomString(5),
		Password: hashedPassword,
	}
	return

}

func TestCreateUserAPI(t *testing.T) {
	user, password := randomUserPassword(t)

	testCases := []struct {
		name          string
		req           *pb.CreateUserRequest
		buildStubs    func(store *mockdb.MockStore, taskDistributor *mockwk.MockTaskDistributor)
		checkResponse func(t *testing.T, res *pb.CreateUserResponse, err error)
	}{{
		name: "OK",
		req: &pb.CreateUserRequest{
			Email:    user.Email,
			Password: password,
			Name:     user.Name,
			LastName: user.LastName,
		},
		buildStubs: func(store *mockdb.MockStore, taskDistributor *mockwk.MockTaskDistributor) {
			arg := db.CreateUserTxParams{
				CreateUserParams: db.CreateUserParams{
					Email:    user.Email,
					Name:     user.Name,
					LastName: user.LastName,
				},
			}
			store.EXPECT().
				CreateUserTx(gomock.Any(), EqCreateUserTxParams(arg, password, user)).
				Times(1).
				Return(db.CreateUserTxResult{User: user}, nil)

			taskPayload := &worker.PayloadSendVerifyEmail{
				Email: user.Email,
			}
			taskDistributor.EXPECT().
				DistributeTaskSendVerifyEmail(gomock.Any(), taskPayload, gomock.Any()).
				Times(1).
				Return(nil)
		},
		checkResponse: func(t *testing.T, res *pb.CreateUserResponse, err error) {
			require.NoError(t, err)
			require.NotNil(t, res)
			createdUser := res.GetUser()
			require.Equal(t, user.Email, createdUser.Email)
			require.Equal(t, user.Name, createdUser.Name)
			require.Equal(t, user.LastName, createdUser.LastName)
		},
	},
		{
			name: "InternalError",
			req: &pb.CreateUserRequest{
				Email:    user.Email,
				Password: password,
				Name:     user.Name,
				LastName: user.LastName,
			},
			buildStubs: func(store *mockdb.MockStore, taskDistributor *mockwk.MockTaskDistributor) {

				store.EXPECT().
					CreateUserTx(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.CreateUserTxResult{}, pgx.ErrTxClosed)

				taskDistributor.EXPECT().
					DistributeTaskSendVerifyEmail(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, res *pb.CreateUserResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.Internal, st.Code())
			},
		},
	}
	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			storeCtrl := gomock.NewController(t)
			defer storeCtrl.Finish()

			store := mockdb.NewMockStore(storeCtrl)

			taskCtrl := gomock.NewController(t)
			defer taskCtrl.Finish()

			taskDistributor := mockwk.NewMockTaskDistributor(taskCtrl)

			tc.buildStubs(store, taskDistributor)

			server := newTestServer(t, store, taskDistributor)
			res, err := server.CreateUser(context.Background(), tc.req)
			tc.checkResponse(t, res, err)
		})

	}
}
