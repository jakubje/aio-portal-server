package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/jakub/aioportal/server/token"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	mockdb "github.com/jakub/aioportal/server/db/mock"
	db "github.com/jakub/aioportal/server/db/sqlc"
	"github.com/jakub/aioportal/server/internal/utils"
	"github.com/jakub/aioportal/server/util"
	"github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

type eqCreateUserParamsMatcher struct {
	arg      db.CreateUserParams
	password string
}

func (e eqCreateUserParamsMatcher) Matches(x interface{}) bool {
	arg, ok := x.(db.CreateUserParams)
	if !ok {
		return false
	}

	err := util.CheckPassword(e.password, arg.Password)
	if err != nil {
		return false
	}

	e.arg.Password = arg.Password
	return reflect.DeepEqual(e.arg, arg)
}

func (e eqCreateUserParamsMatcher) String() string {
	return fmt.Sprintf("matches arg %v and password %v", e.arg, e.password)
}

func EqCreateUserParams(arg db.CreateUserParams, password string) gomock.Matcher {
	return eqCreateUserParamsMatcher{arg, password}
}

func TestCreateUserAPI(t *testing.T) {
	user, password := randomUserPassword(t)

	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recorded *httptest.ResponseRecorder)
	}{{
		name: "OK",
		body: gin.H{
			"email":     user.Email,
			"password":  password,
			"name":      user.Name,
			"last_name": user.LastName,
		},
		buildStubs: func(store *mockdb.MockStore) {
			arg := db.CreateUserParams{
				Email:    user.Email,
				Name:     user.Name,
				LastName: user.LastName,
			}
			store.EXPECT().
				CreateUser(gomock.Any(), EqCreateUserParams(arg, password)).
				Times(1).
				Return(user, nil)
		},
		checkResponse: func(recorded *httptest.ResponseRecorder) {
			require.Equal(t, http.StatusOK, recorded.Code)
			requireBodyMatchAccount(t, recorded.Body, user)
		},
	},
		{
			name: "InternalError",
			body: gin.H{
				"email":     user.Email,
				"password":  password,
				"name":      user.Name,
				"last_name": user.LastName,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.User{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "DuplicateEmail",
			body: gin.H{
				"email":     user.Email,
				"password":  password,
				"name":      user.Name,
				"last_name": user.LastName,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.User{}, &pq.Error{Code: "23505"})
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusForbidden, recorder.Code)
			},
		},
		{
			name: "InvalidEmail",
			body: gin.H{
				"email":     "invalid-email",
				"password":  password,
				"name":      user.Name,
				"last_name": user.LastName,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "TooShortPassword",
			body: gin.H{
				"email":     "invalid-email",
				"name":      user.Name,
				"last_name": user.LastName,
				"password":  "123",
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}
	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			// Marshal body data to JSON
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/users"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)

			tc.checkResponse(recorder)
		})

	}
}

func TestLoginUserAPI(t *testing.T) {
	user, password := randomUserPassword(t)

	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recorded *httptest.ResponseRecorder)
	}{{
		name: "OK",
		body: gin.H{
			"email":    user.Email,
			"password": password,
		},
		buildStubs: func(store *mockdb.MockStore) {
			store.EXPECT().
				GetUser(gomock.Any(), gomock.Eq(user.Email)).
				Times(1).
				Return(user, nil)
			store.EXPECT().
				CreateSession(gomock.Any(), gomock.Any()).
				Times(1)
		},
		checkResponse: func(recorder *httptest.ResponseRecorder) {
			require.Equal(t, http.StatusOK, recorder.Code)
		},
	},
		{
			name: "UserNotFound",
			body: gin.H{
				"email":    "wrongemail@gmail.com",
				"password": password,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.User{}, db.ErrRecordNotFound)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				log.Println(recorder.Body)
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "IncorrectPassword",
			body: gin.H{
				"email":    user.Email,
				"password": "incorrect",
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUser(gomock.Any(), gomock.Eq(user.Email)).
					Times(1).
					Return(user, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "InternalError",
			body: gin.H{
				"email":    user.Email,
				"password": password,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.User{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InvalidEmail",
			body: gin.H{
				"email":     "invalid-email",
				"password":  password,
				"name":      user.Name,
				"last_name": user.LastName,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "TooShortPassword",
			body: gin.H{
				"email":     user.Email,
				"name":      user.Name,
				"last_name": user.LastName,
				"password":  "123",
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}
	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			// Marshal body data to JSON
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/users/login"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)

			tc.checkResponse(recorder)
		})

	}
}

func TestGetUserAPI(t *testing.T) {
	user := randomUser(t)

	testCases := []struct {
		name          string
		body          gin.H
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorded *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute)
			},
			body: gin.H{
				"email": user.Email,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUser(gomock.Any(), gomock.Eq(user.Email)).
					Times(1).
					Return(user, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchAccount(t, recorder.Body, user)
			},
		},
		{
			name: "NotFound",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute)
			},
			body: gin.H{
				"email": user.Email,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUser(gomock.Any(), gomock.Eq(user.Email)).
					Times(1).
					Return(db.User{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorded *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorded.Code)
			},
		},
		{
			name: "InternalError",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute)
			},
			body: gin.H{
				"email": user.Email,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUser(gomock.Any(), gomock.Eq(user.Email)).
					Times(1).
					Return(db.User{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorded *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorded.Code)
			},
		},
		{
			name: "InvalidEmail",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute)
			},
			body: gin.H{
				"email": "invalid-email",
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorded *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorded.Code)
			},
		},
		{
			name: "UnauthorizedUser",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, 123, time.Minute)
			},
			body: gin.H{
				"email": user.Email,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUser(gomock.Any(), gomock.Eq(user.Email)).
					Times(1).
					Return(user, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				log.Println(recorder.Body)
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "NoAuthorization",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
			},
			body: gin.H{
				"email": user.Email,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
	}
	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			// Marshal body data to JSON
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/user"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})

	}
}

//func TestDeleteUserAPI(t *testing.T) {
//	user := randomUser(t)
//
//	testCases := []struct {
//		name          string
//		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
//		buildStubs    func(store *mockdb.MockStore)
//		checkResponse func(t *testing.T, recorded *httptest.ResponseRecorder)
//	}{
//		{
//			name: "OK",
//			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
//				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
//			},
//			buildStubs: func(store *mockdb.MockStore) {
//				store.EXPECT().
//					DeleteUser(gomock.Any(), gomock.Eq(user.Email)).
//					Times(1)
//			},
//			checkResponse: func(t *testing.T, recorded *httptest.ResponseRecorder) {
//				log.Println(recorded.Body)
//				require.Equal(t, http.StatusOK, recorded.Code)
//				requireBodyDeleteAccount(t, recorded.Body, "user deleted")
//			},
//		},
//	}
//	for i := range testCases {
//		tc := testCases[i]
//
//		t.Run(tc.name, func(t *testing.T) {
//			ctrl := gomock.NewController(t)
//			defer ctrl.Finish()
//
//			store := mockdb.NewMockStore(ctrl)
//			tc.buildStubs(store)
//
//			server := newTestServer(t, store)
//			recorder := httptest.NewRecorder()
//
//			url := fmt.Sprintf("/users")
//			request, err := http.NewRequest(http.MethodDelete, url, nil)
//			require.NoError(t, err)
//
//			tc.setupAuth(t, request, server.tokenMaker)
//			server.router.ServeHTTP(recorder, request)
//			tc.checkResponse(t, recorder)
//		})
//
//	}
//}

//func TestListUsersAPI(t *testing.T) {
//
//	n := 5
//	users := make([]db.User, n)
//	for i := 0; i < n; i++ {
//		users[i] = randomUser()
//	}
//
//	type Query struct {
//		pageID   int
//		pageSize int
//	}
//	testCases := []struct {
//		name          string
//		query         Query
//		buildStubs    func(store *mockdb.MockStore)
//		checkResponse func(recorder *httptest.ResponseRecorder)
//		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
//	}{
//		{
//			name: "InternalError",
//
//			query: Query{
//				pageID:   1,
//				pageSize: n,
//			},
//
//			buildStubs: func(store *mockdb.MockStore) {
//				store.EXPECT().
//					ListUsers(gomock.Any(), gomock.Any()).
//					Times(1).
//					Return([]db.User{}, sql.ErrConnDone)
//			},
//			checkResponse: func(recorder *httptest.ResponseRecorder) {
//				require.Equal(t, http.StatusInternalServerError, recorder.Code)
//			},
//		},
//		{
//			name: "InvalidPageID",
//
//			query: Query{
//				pageID:   -1,
//				pageSize: n,
//			},
//
//			buildStubs: func(store *mockdb.MockStore) {
//				store.EXPECT().
//					ListUsers(gomock.Any(), gomock.Any()).
//					Times(0)
//			},
//			checkResponse: func(recorder *httptest.ResponseRecorder) {
//				require.Equal(t, http.StatusBadRequest, recorder.Code)
//			},
//		},
//		{
//			name: "InvalidPageSize",
//			query: Query{
//				pageID:   1,
//				pageSize: 100000,
//			},
//			buildStubs: func(store *mockdb.MockStore) {
//				store.EXPECT().
//					ListUsers(gomock.Any(), gomock.Any()).
//					Times(0)
//			},
//			checkResponse: func(recorder *httptest.ResponseRecorder) {
//				require.Equal(t, http.StatusBadRequest, recorder.Code)
//			},
//		},
//	}
//	for i := range testCases {
//		tc := testCases[i]
//
//		t.Run(tc.name, func(t *testing.T) {
//			ctrl := gomock.NewController(t)
//			defer ctrl.Finish()
//
//			store := mockdb.NewMockStore(ctrl)
//			tc.buildStubs(store)
//
//			server := newTestServer(t, store)
//			recorder := httptest.NewRecorder()
//
//			url := "/users"
//			request, err := http.NewRequest(http.MethodGet, url, nil)
//			require.NoError(t, err)
//
//			// Add query parameters to request URL
//			q := request.URL.Query()
//			q.Add("page_id", fmt.Sprintf("%d", tc.query.pageID))
//			q.Add("page_size", fmt.Sprintf("%d", tc.query.pageSize))
//			request.URL.RawQuery = q.Encode()
//
//			tc.setupAuth(t, request, server.tokenMaker)
//			server.router.ServeHTTP(recorder, request)
//			tc.checkResponse(recorder)
//		})
//	}
//}
//
//func TestUpdateUserAPI(t *testing.T) {
//	user := randomUser(t)
//	newName := utils.RandomString(6)
//	newLastName := utils.RandomString(6)
//	newEmail := utils.RandomEmail()
//
//	hashedPassword, err := utils.HashPasswod(utils.RandomString(6))
//	require.NoError(t, err)
//
//	testCases := []struct {
//		id            int64
//		name          string
//		body          gin.H
//		buildStubs    func(store *mockdb.MockStore)
//		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
//		checkResponse func(recorded *httptest.ResponseRecorder, res *db.UpdateUserParams, err error)
//	}{{
//		name: "OK",
//		setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
//			addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute)
//		},
//		body: gin.H{
//			"email":     utils.RandomEmail(),
//			"password":  hashedPassword,
//			"name":      utils.RandomString(9),
//			"last_name": utils.RandomString(9),
//		},
//		buildStubs: func(store *mockdb.MockStore) {
//
//			arg := db.UpdateUserParams{
//				Email:    newEmail,
//				Name:     newName,
//				LastName: newLastName,
//				Password: hashedPassword,
//			}
//			updatedUser := db.User{
//				Email:             newEmail,
//				Name:              newName,
//				LastName:          newLastName,
//				Password:          user.Password,
//				PasswordChangedAt: user.PasswordChangedAt,
//			}
//			store.EXPECT().
//				UpdateUser(gomock.Any(), gomock.Eq(arg)).
//				Times(1).
//				Return(updatedUser, nil)
//		},
//		checkResponse: func(recorded *httptest.ResponseRecorder, res *db.UpdateUserParams, err error) {
//			log.Println(recorded.Body)
//			require.NoError(t, err)
//			require.NotNil(t, res)
//
//			updatedUser := res
//			require.Equal(t, newEmail, updatedUser.Email)
//			require.Equal(t, newName, updatedUser.Name)
//			require.Equal(t, newLastName, updatedUser.LastName)
//		},
//	},
//	//{
//	//	name: "InternalError",
//	//	setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
//	//		addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute)
//	//	},
//	//	body: gin.H{
//	//		"email":     utils.RandomEmail(),
//	//		"password":  hashedPassword,
//	//		"name":      utils.RandomString(9),
//	//		"last_name": utils.RandomString(9),
//	//	},
//	//	buildStubs: func(store *mockdb.MockStore) {
//	//		store.EXPECT().
//	//			UpdateUser(gomock.Any(), gomock.Any()).
//	//			Times(1).
//	//			Return(db.User{}, sql.ErrConnDone)
//	//	},
//	//	checkResponse: func(recorder *httptest.ResponseRecorder) {
//	//		require.Equal(t, http.StatusInternalServerError, recorder.Code)
//	//	},
//	//},
//	//{
//	//	name: "DuplicateEmail",
//	//	body: gin.H{
//	//		"email":     user.Email,
//	//		"password":  hashedPassword,
//	//		"name":      user.Name,
//	//		"last_name": user.LastName,
//	//	},
//	//	buildStubs: func(store *mockdb.MockStore) {
//	//		store.EXPECT().
//	//			CreateUser(gomock.Any(), gomock.Any()).
//	//			Times(1).
//	//			Return(db.User{}, &pq.Error{Code: "23505"})
//	//	},
//	//	checkResponse: func(recorder *httptest.ResponseRecorder) {
//	//		require.Equal(t, http.StatusForbidden, recorder.Code)
//	//	},
//	//},
//	//{
//	//	name: "InvalidEmail",
//	//	setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
//	//		addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute)
//	//	},
//	//	body: gin.H{
//	//		"email":     "invalid-email",
//	//		"password":  hashedPassword,
//	//		"name":      user.Name,
//	//		"last_name": user.LastName,
//	//	},
//	//	buildStubs: func(store *mockdb.MockStore) {
//	//		store.EXPECT().
//	//			CreateUser(gomock.Any(), gomock.Any()).
//	//			Times(0)
//	//	},
//	//	checkResponse: func(recorder *httptest.ResponseRecorder) {
//	//		require.Equal(t, http.StatusBadRequest, recorder.Code)
//	//	},
//	//},
//	//{
//	//	name: "TooShortPassword",
//	//	setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
//	//		addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute)
//	//	},
//	//	body: gin.H{
//	//		"email":     "invalid-email",
//	//		"name":      user.Name,
//	//		"last_name": user.LastName,
//	//		"password":  "123",
//	//	},
//	//	buildStubs: func(store *mockdb.MockStore) {
//	//		store.EXPECT().
//	//			CreateUser(gomock.Any(), gomock.Any()).
//	//			Times(0)
//	//	},
//	//	checkResponse: func(recorder *httptest.ResponseRecorder) {
//	//		require.Equal(t, http.StatusBadRequest, recorder.Code)
//	//	},
//	//},
//	}
//	for i := range testCases {
//		tc := testCases[i]
//
//		t.Run(tc.name, func(t *testing.T) {
//
//			ctrl := gomock.NewController(t)
//			defer ctrl.Finish()
//
//			store := mockdb.NewMockStore(ctrl)
//			tc.buildStubs(store)
//
//			server := newTestServer(t, store)
//			recorder := httptest.NewRecorder()
//
//			// Marshal body data to JSON
//			data, err := json.Marshal(tc.body)
//			require.NoError(t, err)
//
//			url := "/users/update"
//			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
//			require.NoError(t, err)
//
//			tc.setupAuth(t, request, server.tokenMaker)
//
//			server.router.ServeHTTP(recorder, request)
//			tc.checkResponse(recorder)
//		})
//
//	}
//}

func randomUser(t *testing.T) (user db.User) {
	password := utils.RandomString(8)
	hashedPassword, err := utils.HashPasswod(password)
	require.NoError(t, err)

	user = db.User{
		Email:    utils.RandomEmail(),
		Name:     utils.RandomString(6),
		LastName: utils.RandomString(5),
		Password: hashedPassword,
	}
	return

}

func randomUserPassword(t *testing.T) (user db.User, password string) {
	password = utils.RandomString(8)
	hashedPassword, err := utils.HashPasswod(password)
	require.NoError(t, err)

	user = db.User{
		Email:    utils.RandomEmail(),
		Name:     utils.RandomString(6),
		LastName: utils.RandomString(5),
		Password: hashedPassword,
	}
	return

}

func requireBodyMatchAccount(t *testing.T, body *bytes.Buffer, user db.User) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotUser db.User
	err = json.Unmarshal(data, &gotUser)
	require.NoError(t, err)
	require.Equal(t, user.Name, gotUser.Name)
	require.Equal(t, user.LastName, gotUser.LastName)
	require.Equal(t, user.Email, gotUser.Email)
	require.Empty(t, gotUser.Password)
}

func requireBodyDeleteAccount(t *testing.T, body *bytes.Buffer, message string) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	type deleteMessage struct {
		Message string `json:"message"`
	}
	var deleteResponse deleteMessage

	err = json.Unmarshal(data, &deleteResponse)
	require.NoError(t, err)
	require.Equal(t, message, deleteResponse.Message)
}
