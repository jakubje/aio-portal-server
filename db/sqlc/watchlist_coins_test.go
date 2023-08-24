package db

import (
	"context"

	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomWatchlistCoin(t *testing.T) WatchlistCoin {
	coin := CreateRandomCoin(t)
	_, watchlist := CreateRandomWatchlist(t)

	arg := AddWatchlistCoinParams{
		WatchlistID: watchlist.ID,
		CoinID:      coin.CoinID,
	}
	watchlistCoin, err := testStore.AddWatchlistCoin(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, watchlistCoin)

	require.Equal(t, arg.WatchlistID, watchlistCoin.WatchlistID)
	require.Equal(t, arg.CoinID, watchlistCoin.CoinID)

	return watchlistCoin
}

func addCoinToWatchlist(t *testing.T, watchlist *Watchlist) WatchlistCoin {

	coin := CreateRandomCoin(t)
	arg := AddWatchlistCoinParams{
		WatchlistID: watchlist.ID,
		CoinID:      coin.CoinID,
	}
	watchlistCoin, err := testStore.AddWatchlistCoin(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, watchlistCoin)

	require.Equal(t, arg.WatchlistID, watchlistCoin.WatchlistID)
	require.Equal(t, arg.CoinID, watchlistCoin.CoinID)

	return watchlistCoin
}

func TestCreateWatchlistCoin(t *testing.T) {
	createRandomWatchlistCoin(t)
}

func TestDeleteWatchlistCoin(t *testing.T) {
	coin1 := createRandomWatchlistCoin(t)

	arg := RemoveWatchlistCoinParams{
		WatchlistID: coin1.WatchlistID,
		CoinID:      coin1.CoinID,
	}
	err := testStore.RemoveWatchlistCoin(context.Background(), arg)
	require.NoError(t, err)
}

func TestListWatchlistCoins(t *testing.T) {
	user, watchList := CreateRandomWatchlist(t)

	for i := 0; i < 10; i++ {
		addCoinToWatchlist(t, &watchList)
	}

	arg := ListWatchlistCoinsParams{
		WatchlistID: watchList.ID,
		AccountID:   user.ID,
	}
	
	watchlistCoins, err := testStore.ListWatchlistCoins(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, watchlistCoins, 10)

	for _, watchlistCoin := range watchlistCoins {
		require.NotEmpty(t, watchlistCoin)
	}
}
