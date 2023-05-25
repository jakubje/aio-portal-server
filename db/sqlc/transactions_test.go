package db

import (
	"context"
	"database/sql"
	"server/internal/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomTransaction(t *testing.T) Transaction {
	user, portfolio := createUserAndPortfolio(t)
	coin := createRandomPortfolioCoin(t)
	arg := CreateTransactionParams{
		AccountID:      user.ID,
		PortfolioID:    portfolio.ID,
		CoinID:         coin.ID,
		Symbol:         coin.CoinSymbol,
		Type:           0,
		Quantity:       float64(utils.RandomInt()),
		PricePerCoin:   float64(utils.RandomInt()),
		TimeTransacted: time.Time{},
	}

	transaction, err := testQueries.CreateTransaction(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transaction)

	require.Equal(t, arg.AccountID, transaction.AccountID)
	require.Equal(t, arg.PortfolioID, transaction.PortfolioID)
	require.Equal(t, arg.CoinID, transaction.CoinID)
	require.Equal(t, arg.Symbol, transaction.Symbol)
	require.Equal(t, arg.Type, transaction.Type)
	require.Equal(t, arg.Quantity, transaction.Quantity)
	require.Equal(t, arg.PricePerCoin, transaction.PricePerCoin)
	require.WithinDuration(t, arg.TimeTransacted, transaction.TimeTransacted, time.Second)

	return transaction
}

func TestCreateTransaction(t *testing.T) {
	createRandomTransaction(t)
}

func TestGetTransaction(t *testing.T) {
	transaction1 := createRandomTransaction(t)
	transaction2, err := testQueries.GetTransaction(context.Background(), transaction1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transaction2)

	require.Equal(t, transaction1.AccountID, transaction2.AccountID)
	require.Equal(t, transaction1.PortfolioID, transaction2.PortfolioID)
	require.Equal(t, transaction1.CoinID, transaction2.CoinID)
	require.Equal(t, transaction1.Symbol, transaction2.Symbol)
	require.Equal(t, transaction1.Type, transaction2.Type)
	require.Equal(t, transaction1.Quantity, transaction2.Quantity)
	require.Equal(t, transaction1.PricePerCoin, transaction2.PricePerCoin)
	require.WithinDuration(t, transaction1.TimeTransacted, transaction2.TimeTransacted, time.Second)
	require.WithinDuration(t, transaction1.TimeCreated, transaction2.TimeCreated, time.Second)
}

func TestDeleteTransaction(t *testing.T) {
	transaction1 := createRandomTransaction(t)
	err := testQueries.DeleteTransaction(context.Background(), transaction1.ID)
	require.NoError(t, err)

	transaction2, err := testQueries.GetTransaction(context.Background(), transaction1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, transaction2)
}

func createTransactionsForAccount(t *testing.T, user *User, portfolio *Portfolio) Transaction {
	coin := createRandomPortfolioCoin(t)

	arg := CreateTransactionParams{
		PortfolioID:    portfolio.ID,
		AccountID:      user.ID,
		CoinID:         coin.ID,
		Symbol:         coin.CoinSymbol,
		Type:           0,
		Quantity:       float64(utils.RandomInt()),
		PricePerCoin:   float64(utils.RandomInt()),
		TimeTransacted: time.Time{},
	}

	transaction, err := testQueries.CreateTransaction(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transaction)

	require.Equal(t, arg.AccountID, transaction.AccountID)
	require.Equal(t, arg.PortfolioID, transaction.PortfolioID)
	require.Equal(t, arg.CoinID, transaction.CoinID)
	require.Equal(t, arg.Symbol, transaction.Symbol)
	require.Equal(t, arg.Type, transaction.Type)
	require.Equal(t, arg.Quantity, transaction.Quantity)
	require.Equal(t, arg.PricePerCoin, transaction.PricePerCoin)
	require.WithinDuration(t, arg.TimeTransacted, transaction.TimeTransacted, time.Second)

	return transaction
}

func TestListTransactionsByAccount(t *testing.T) {
	user, portfolio := createUserAndPortfolio(t)

	for i := 0; i < 10; i++ {
		createTransactionsForAccount(t, &user, &portfolio)
	}

	arg := ListTransactionsByAccountParams{
		AccountID: user.ID,
		Limit:     10,
		Offset:    0,
	}
	userTransactions, err := testQueries.ListTransactionsByAccount(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, userTransactions, 10)

	for _, userTransaction := range userTransactions {
		require.NotEmpty(t, userTransaction)
	}

}

func createTransactionsForAccountForCoin(t *testing.T, user *User, portfolio *Portfolio, coin *Coin) Transaction {

	arg := CreateTransactionParams{
		AccountID:      user.ID,
		PortfolioID:    portfolio.ID,
		CoinID:         coin.ID,
		Symbol:         coin.CoinSymbol,
		Type:           0,
		Quantity:       float64(utils.RandomInt()),
		PricePerCoin:   float64(utils.RandomInt()),
		TimeTransacted: time.Time{},
	}

	transaction, err := testQueries.CreateTransaction(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transaction)

	require.Equal(t, arg.AccountID, transaction.AccountID)
	require.Equal(t, arg.PortfolioID, transaction.PortfolioID)
	require.Equal(t, arg.CoinID, transaction.CoinID)
	require.Equal(t, arg.Symbol, transaction.Symbol)
	require.Equal(t, arg.Type, transaction.Type)
	require.Equal(t, arg.Quantity, transaction.Quantity)
	require.Equal(t, arg.PricePerCoin, transaction.PricePerCoin)
	require.WithinDuration(t, arg.TimeTransacted, transaction.TimeTransacted, time.Second)

	return transaction
}

func TestListTransactionsByAccountByCoin(t *testing.T) {
	user, portfolio := createUserAndPortfolio(t)
	coin := createRandomPortfolioCoin(t)

	arg := ListTransactionsByAccountByCoinParams{
		AccountID: user.ID,
		CoinID:    coin.ID,
		Limit:     10,
		Offset:    0,
	}

	for i := 0; i < 10; i++ {
		createTransactionsForAccountForCoin(t, &user, &portfolio, &coin)
	}

	userTransactions, err := testQueries.ListTransactionsByAccountByCoin(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, userTransactions, 10)

	for _, userTransaction := range userTransactions {
		require.NotEmpty(t, userTransaction)
	}
}

func TestGetRollUpByCoinByPortfolio(t *testing.T) {
	user, portfolio := createUserAndPortfolio(t)

	for i := 0; i < 10; i++ {
		coin := createRandomPortfolioCoin(t)
		createTransactionsForAccountForCoin(t, &user, &portfolio, &coin)
	}

	rollUp, err := testQueries.GetRollUpByCoinByPortfolio(context.Background(), portfolio.ID)

	require.NoError(t, err)
	require.Len(t, rollUp, 10)

	for _, rollUp := range rollUp {
		require.NotEmpty(t, rollUp)
	}
}
