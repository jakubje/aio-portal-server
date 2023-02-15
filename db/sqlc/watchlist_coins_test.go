package db

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"server/internal/utils"
	"testing"
)

func createRandomWatchlistCoin(t *testing.T) WatchlistCoin {
	watchlist := CreateRandomWatchlist(t)
	arg := CreateWatchlistCoinsParams{
		WatchlistID: watchlist.ID,
		Name:        utils.RandomString(5),
		Symbol:      utils.RandomString(3),
		Rank:        int16(utils.RandomInt()),
	}
	watchlistCoin, err := testQueries.CreateWatchlistCoins(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, watchlistCoin)

	require.Equal(t, arg.WatchlistID, watchlistCoin.WatchlistID)
	require.Equal(t, arg.Name, watchlistCoin.Name)
	require.Equal(t, arg.Symbol, watchlistCoin.Symbol)
	require.Equal(t, arg.Rank, watchlistCoin.Rank)

	require.NotZero(t, watchlistCoin.ID)

	return watchlistCoin
}

func createWatchlistCoinsForWatchlist(t *testing.T, watchlist *Watchlist) WatchlistCoin {

	arg := CreateWatchlistCoinsParams{
		WatchlistID: watchlist.ID,
		Name:        utils.RandomString(5),
		Symbol:      utils.RandomString(3),
		Rank:        int16(utils.RandomInt()),
	}
	watchlistCoin, err := testQueries.CreateWatchlistCoins(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, watchlistCoin)

	require.Equal(t, arg.WatchlistID, watchlistCoin.WatchlistID)
	require.Equal(t, arg.Name, watchlistCoin.Name)
	require.Equal(t, arg.Symbol, watchlistCoin.Symbol)
	require.Equal(t, arg.Rank, watchlistCoin.Rank)

	require.NotZero(t, watchlistCoin.ID)

	return watchlistCoin
}

func TestCreateWatchlistCoin(t *testing.T) {
	createRandomWatchlistCoin(t)
}

func TestGetWachlistCoin(t *testing.T) {
	coin1 := createRandomWatchlistCoin(t)
	coin2, err := testQueries.GetWatchlistCoin(context.Background(), coin1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, coin2)

	require.Equal(t, coin1.ID, coin2.ID)
	require.Equal(t, coin1.Name, coin2.Name)
	require.Equal(t, coin1.Symbol, coin2.Symbol)
	require.Equal(t, coin1.Rank, coin2.Rank)
	require.Equal(t, coin1.WatchlistID, coin2.WatchlistID)
}

func TestDeleteWatchlistCoin(t *testing.T) {
	coin1 := createRandomWatchlistCoin(t)
	err := testQueries.DeleteWatchlistCoin(context.Background(), coin1.ID)
	require.NoError(t, err)

	coin2, err := testQueries.GetWatchlistCoin(context.Background(), coin1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, coin2)
}

func TestListWatchlistCoins(t *testing.T) {
	watchList := CreateRandomWatchlist(t)

	for i := 0; i < 10; i++ {
		createWatchlistCoinsForWatchlist(t, &watchList)
	}

	watchlistCoins, err := testQueries.ListWatchlistsCoins(context.Background(), watchList.ID)
	require.NoError(t, err)
	require.Len(t, watchlistCoins, 10)

	for _, watchlistCoin := range watchlistCoins {
		require.NotEmpty(t, watchlistCoin)
	}
}
