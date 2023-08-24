package api

import (
	"bytes"
	"encoding/json"
	"io"
	"testing"

	db "github.com/jakub/aioportal/server/db/sqlc"
	"github.com/jakub/aioportal/server/util"
	"github.com/stretchr/testify/require"
)

//
//func TestCreateFootballAPI(t *testing.T) {
//	user := randomUser(t)
//	football := createRandomFootball(user)
//	log.Println(football)
//
//	testCases := []struct {
//		name          string
//		body          gin.H
//		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
//		buildStubs    func(store *mockdb.MockStore)
//		checkResponse func(t *testing.T, recorded *httptest.ResponseRecorder)
//	}{{
//		name: "OK",
//		body: gin.H{
//			"team":    football.Team,
//			"league":  football.League,
//			"country": football.Country,
//		},
//		setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
//			addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute)
//		},
//		buildStubs: func(store *mockdb.MockStore) {
//			arg := db.CreateFootballParams{
//				Team:    football.Team,
//				League:  football.League,
//				Country: football.Country,
//			}
//			store.EXPECT().
//				CreateFootball(gomock.Any(), gomock.Eq(arg)).
//				Times(1).
//				Return(football, nil)
//		},
//		checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
//			require.Equal(t, http.StatusOK, recorder.Code)
//			requireBodyMatchFootball(t, recorder.Body, football)
//		},
//	}}
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
//			// Marshal body data to JSON
//			data, err := json.Marshal(tc.body)
//			require.NoError(t, err)
//
//			url := "/football"
//			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
//			require.NoError(t, err)
//
//			tc.setupAuth(t, request, server.tokenMaker)
//			server.router.ServeHTTP(recorder, request)
//			tc.checkResponse(t, recorder)
//		})
//
//	}
//
//}

func createRandomFootball(user db.User) (football db.Football) {
	team := util.RandomString(6)
	league := util.RandomString(6)
	country := util.RandomString(6)

	football = db.Football{
		AccountID: user.ID,
		Team:      team,
		League:    league,
		Country:   country,
	}
	return
}

func requireBodyMatchFootball(t *testing.T, body *bytes.Buffer, football db.Football) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotFootball db.Football
	err = json.Unmarshal(data, &gotFootball)
	require.NoError(t, err)
	require.Equal(t, football.Team, gotFootball.Team)
	require.Equal(t, football.League, gotFootball.League)
	require.Equal(t, football.Country, gotFootball.Country)
}
