package db

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"server/internal/utils"
	"strconv"
	"testing"
	"time"
)

func createRandomPortfolioCoin(t *testing.T) Coin {
	portfolio := createRandomPortfolio(t)
	arg := CreateCoinParams{
		PortfolioID: portfolio.ID,
		CoinName:    utils.RandomString(5),
		CoinSymbol:  utils.RandomString(3),
		Amount:      int32(utils.RandomInt()),
		NoOfCoins:   strconv.Itoa(int(utils.RandomInt())),
	}

	coin, err := testQueries.CreateCoin(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, coin)

	require.Equal(t, arg.PortfolioID, coin.PortfolioID)
	require.Equal(t, arg.CoinName, coin.CoinName)
	require.Equal(t, arg.CoinSymbol, coin.CoinSymbol)
	require.Equal(t, arg.Amount, coin.Amount)
	require.Equal(t, arg.NoOfCoins, coin.NoOfCoins)
	require.NotZero(t, coin.ID)

	return coin

}

func TestCreateCoin(t *testing.T) {
	createRandomPortfolioCoin(t)
}

func TestGetCoin(t *testing.T) {
	coin1 := createRandomPortfolioCoin(t)
	coin2, err := testQueries.GetCoin(context.Background(), coin1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, coin2)

	require.Equal(t, coin1.ID, coin2.ID)
	require.Equal(t, coin1.CoinSymbol, coin2.CoinSymbol)
	require.Equal(t, coin1.PortfolioID, coin2.PortfolioID)
	require.Equal(t, coin1.NoOfCoins, coin2.NoOfCoins)
	require.Equal(t, coin1.Amount, coin2.Amount)
	require.Equal(t, coin1.CoinName, coin2.CoinName)
	require.WithinDuration(t, coin1.TimeCreated, coin2.TimeCreated, time.Second)
}

func TestUpdateCoin(t *testing.T) {
	coin1 := createRandomPortfolioCoin(t)

	arg := UpdateCoinParams{
		ID:        coin1.ID,
		Amount:    int32(utils.RandomInt()),
		NoOfCoins: strconv.Itoa(int(utils.RandomInt())),
	}

	coin2, err := testQueries.UpdateCoin(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, coin2)

	require.Equal(t, arg.ID, coin2.ID)
	require.Equal(t, arg.Amount, coin2.Amount)
	require.Equal(t, arg.NoOfCoins, coin2.NoOfCoins)

}

func TestDeleteCoin(t *testing.T) {
	coin1 := createRandomPortfolioCoin(t)
	err := testQueries.DeleteCoin(context.Background(), coin1.ID)
	require.NoError(t, err)

	coin2, err := testQueries.GetCoin(context.Background(), coin1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, coin2)
}

func createCoinByPortfolioId(t *testing.T, portfolio *Portfolio) Coin {
	arg := CreateCoinParams{
		PortfolioID: portfolio.ID,
		CoinName:    utils.RandomString(5),
		CoinSymbol:  utils.RandomString(3),
		Amount:      int32(utils.RandomInt()),
		NoOfCoins:   strconv.Itoa(int(utils.RandomInt())),
	}

	coin, err := testQueries.CreateCoin(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, coin)

	require.Equal(t, arg.PortfolioID, coin.PortfolioID)
	require.Equal(t, arg.CoinName, coin.CoinName)
	require.Equal(t, arg.CoinSymbol, coin.CoinSymbol)
	require.Equal(t, arg.Amount, coin.Amount)
	require.Equal(t, arg.NoOfCoins, coin.NoOfCoins)

	return coin

}

func TestListCoinsByPortfolioId(t *testing.T) {
	portfolio := createRandomPortfolio(t)

	for i := 0; i < 10; i++ {
		createCoinByPortfolioId(t, &portfolio)
	}

	portfolioCoins, err := testQueries.ListCoins(context.Background(), portfolio.ID)
	require.NoError(t, err)
	require.Len(t, portfolioCoins, 10)

	for _, portfolioCoin := range portfolioCoins {
		require.NotEmpty(t, portfolioCoin)
	}
}
