package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/golang/mock/gomock"
	mockdb "github.com/jakub/aioportal/server/db/mock"
	db "github.com/jakub/aioportal/server/db/sqlc"
	"github.com/jakub/aioportal/server/internal/utils"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUserAPI(t *testing.T) {
	user := randomUser()

	testCases := []struct {
		name          string
		accountID     int64
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorded *httptest.ResponseRecorder)
	}{
		{
			name:      "OK",
			accountID: user.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUser(gomock.Any(), gomock.Eq(user.ID)).
					Times(1).
					Return(user, nil)
			},
			checkResponse: func(t *testing.T, recorded *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorded.Code)
				requireBodyMatchAccount(t, recorded.Body, user)
			},
		},
		{
			name:      "NotFound",
			accountID: user.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUser(gomock.Any(), gomock.Eq(user.ID)).
					Times(1).
					Return(db.User{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorded *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorded.Code)
			},
		},
		{
			name:      "InternalError",
			accountID: user.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUser(gomock.Any(), gomock.Eq(user.ID)).
					Times(1).
					Return(db.User{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorded *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorded.Code)
			},
		},
		{
			name:      "InvalidID",
			accountID: 0,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorded *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorded.Code)
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

			server := NewServer(store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/user/%d", tc.accountID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})

	}
}

func TestDeleteUserAPI(t *testing.T) {
	user := randomUser()

	testCases := []struct {
		name          string
		accountID     int64
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorded *httptest.ResponseRecorder)
	}{
		{
			name:      "OK",
			accountID: user.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					DeleteUser(gomock.Any(), gomock.Eq(user.ID)).
					Times(1)
			},
			checkResponse: func(t *testing.T, recorded *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorded.Code)
				requireBodyDeleteAccount(t, recorded.Body, "user deleted")
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

			server := NewServer(store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/user/%d", tc.accountID)
			request, err := http.NewRequest(http.MethodDelete, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})

	}
}

func TestListUsersAPI(t *testing.T) {

	n := 5
	users := make([]db.User, n)
	for i := 0; i < n; i++ {
		users[i] = randomUser()
	}

	type Query struct {
		pageID   int
		pageSize int
	}
	testCases := []struct {
		name          string
		query         Query
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "InternalError",
			query: Query{
				pageID:   1,
				pageSize: n,
			},

			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListUsers(gomock.Any(), gomock.Any()).
					Times(1).
					Return([]db.User{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InvalidPageID",
			query: Query{
				pageID:   -1,
				pageSize: n,
			},

			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListUsers(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "InvalidPageSize",
			query: Query{
				pageID:   1,
				pageSize: 100000,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListUsers(gomock.Any(), gomock.Any()).
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

			server := NewServer(store)
			recorder := httptest.NewRecorder()

			url := "/users"
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			// Add query parameters to request URL
			q := request.URL.Query()
			q.Add("page_id", fmt.Sprintf("%d", tc.query.pageID))
			q.Add("page_size", fmt.Sprintf("%d", tc.query.pageSize))
			request.URL.RawQuery = q.Encode()

			//tc.setupAuth(t, request, server.tokenMaker)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func randomUser() db.User {
	return db.User{
		ID:       int64(utils.RandomInt()),
		Email:    utils.RandomEmail(),
		Name:     utils.RandomString(5),
		LastName: utils.RandomString(5),
		Password: utils.RandomString(8),
	}
}

func requireBodyMatchAccount(t *testing.T, body *bytes.Buffer, user db.User) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotUser db.User
	err = json.Unmarshal(data, &gotUser)
	require.NoError(t, err)
	require.Equal(t, user, gotUser)
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
