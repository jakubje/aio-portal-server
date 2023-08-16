package db

import (
	"context"

	"testing"

	"github.com/jakub/aioportal/server/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomWatchlist(t *testing.T) (User, Watchlist) {
	user := createRandomUser(t)
	arg := CreateWatchlistParams{
		Name:      util.RandomString(5),
		AccountID: user.ID,
	}
	watchlist, err := testStore.CreateWatchlist(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, watchlist)

	require.Equal(t, arg.Name, watchlist.Name)
	require.Equal(t, arg.AccountID, user.ID)

	return user, watchlist
}

func TestCreateWatchlist(t *testing.T) {
	CreateRandomWatchlist(t)
}

func TestGetWatchlist(t *testing.T) {
	user, watchlist1 := CreateRandomWatchlist(t)
	arg := GetWatchlistParams{
		ID:        watchlist1.ID,
		AccountID: user.ID,
	}
	watchlist2, err := testStore.GetWatchlist(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, watchlist2)

	require.Equal(t, watchlist1.ID, watchlist2.ID)
	require.Equal(t, watchlist1.Name, watchlist2.Name)
	require.Equal(t, watchlist1.AccountID, watchlist2.AccountID)
}

func TestUpdateWatchlist(t *testing.T) {
	user, watchlist1 := CreateRandomWatchlist(t)
	arg := UpdateWatchlistParams{
		ID:        watchlist1.ID,
		Name:      util.RandomString(6),
		AccountID: user.ID,
	}
	watchlist2, err := testStore.UpdateWatchlist(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, watchlist2)

	require.Equal(t, watchlist1.AccountID, watchlist2.AccountID)
	require.Equal(t, arg.ID, watchlist2.ID)
	require.Equal(t, arg.Name, watchlist2.Name)
}

func TestDeleteWatchlist(t *testing.T) {
	user, watchlist1 := CreateRandomWatchlist(t)
	arg := DeleteWatchlistParams{
		ID:        watchlist1.ID,
		AccountID: user.ID,
	}
	err := testStore.DeleteWatchlist(context.Background(), arg)
	require.NoError(t, err)

	watchlist2, err := testStore.GetWatchlist(context.Background(), GetWatchlistParams(arg))
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, watchlist2)
}

func createWatchlistForUser(t *testing.T, user *User) Watchlist {

	arg := CreateWatchlistParams{
		Name:      util.RandomString(5),
		AccountID: user.ID,
	}
	watchlist, err := testStore.CreateWatchlist(context.Background(), arg)
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

	watchlists, err := testStore.ListWatchlists(context.Background(), user.ID)
	require.NoError(t, err)
	require.Len(t, watchlists, 10)
	for _, watchlist := range watchlists {
		require.NotEmpty(t, watchlist)
	}

}
