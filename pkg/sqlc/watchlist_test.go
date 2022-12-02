package db

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"server/internal/utils"
	"testing"
)

func CreateRandomWatchlist(t *testing.T) Watchlist {
	user := createRandomUser(t)
	arg := CreateWatchlistParams{
		Name:      utils.RandomString(5),
		AccountID: user.ID,
	}
	watchlist, err := testQueries.CreateWatchlist(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, watchlist)

	require.Equal(t, arg.Name, watchlist.Name)
	require.Equal(t, arg.AccountID, user.ID)

	return watchlist
}

func TestCreateWatchlist(t *testing.T) {
	CreateRandomWatchlist(t)
}

func TestGetWatchlist(t *testing.T) {
	watchlist1 := CreateRandomWatchlist(t)
	watchlist2, err := testQueries.GetWatchlist(context.Background(), watchlist1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, watchlist2)

	require.Equal(t, watchlist1.ID, watchlist2.ID)
	require.Equal(t, watchlist1.Name, watchlist2.Name)
	require.Equal(t, watchlist1.AccountID, watchlist2.AccountID)
}

func TestUpdateWatchlist(t *testing.T) {
	watchlist1 := CreateRandomWatchlist(t)
	arg := UpdateWatchlistParams{
		ID:   watchlist1.ID,
		Name: utils.RandomString(6),
	}
	watchlist2, err := testQueries.UpdateWatchlist(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, watchlist2)

	require.Equal(t, watchlist1.AccountID, watchlist2.AccountID)
	require.Equal(t, arg.ID, watchlist2.ID)
	require.Equal(t, arg.Name, watchlist2.Name)
}

func TestDeleteWatchlist(t *testing.T) {
	watchlist1 := CreateRandomWatchlist(t)
	err := testQueries.DeleteWatchlist(context.Background(), watchlist1.ID)
	require.NoError(t, err)

	watchlist2, err := testQueries.GetWatchlist(context.Background(), watchlist1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, watchlist2)
}

func createWatchlistForUser(t *testing.T, user *User) Watchlist {

	arg := CreateWatchlistParams{
		Name:      utils.RandomString(5),
		AccountID: user.ID,
	}
	watchlist, err := testQueries.CreateWatchlist(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, watchlist)

	require.Equal(t, arg.Name, watchlist.Name)
	require.Equal(t, arg.AccountID, user.ID)

	return watchlist
}

func TestListWatchlists(t *testing.T) {
	user := createRandomUser(t)
	for i := 0; i < 10; i++ {
		createWatchlistForUser(t, &user)
	}

	watchlists, err := testQueries.ListWatchlists(context.Background(), user.ID)
	require.NoError(t, err)
	require.Len(t, watchlists, 10)
	for _, watchlist := range watchlists {
		require.NotEmpty(t, watchlist)
	}

}
