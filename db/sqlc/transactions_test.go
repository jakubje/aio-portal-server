package db

import (
	"context"

	"testing"
	"time"

	"github.com/jakub/aioportal/server/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransaction(t *testing.T) (User, Transaction) {
	user, portfolio := createUserAndPortfolio(t)
	coin := util.RandomString(3)
	arg := CreateTransactionParams{
		AccountID:      user.ID,
		PortfolioID:    portfolio.ID,
		Symbol:         coin,
		Type:           0,
		Quantity:       float64(util.RandomInt()),
		PricePerCoin:   float64(util.RandomInt()),
		TimeTransacted: time.Time{},
	}

	transaction, err := testStore.CreateTransaction(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transaction)

	require.Equal(t, arg.AccountID, transaction.AccountID)
	require.Equal(t, arg.PortfolioID, transaction.PortfolioID)
	require.Equal(t, arg.Symbol, transaction.Symbol)
	require.Equal(t, arg.Type, transaction.Type)
	require.Equal(t, arg.Quantity, transaction.Quantity)
	require.Equal(t, arg.PricePerCoin, transaction.PricePerCoin)
	require.WithinDuration(t, arg.TimeTransacted, transaction.TimeTransacted, time.Second)

	return user, transaction
}

func TestCreateTransaction(t *testing.T) {
	createRandomTransaction(t)
}

func TestGetTransaction(t *testing.T) {
	user, transaction1 := createRandomTransaction(t)
	arg := GetTransactionParams{
		ID:        transaction1.ID,
		AccountID: user.ID,
	}
	transaction2, err := testStore.GetTransaction(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transaction2)

	require.Equal(t, transaction1.AccountID, transaction2.AccountID)
	require.Equal(t, transaction1.PortfolioID, transaction2.PortfolioID)
	require.Equal(t, transaction1.Symbol, transaction2.Symbol)
	require.Equal(t, transaction1.Type, transaction2.Type)
	require.Equal(t, transaction1.Quantity, transaction2.Quantity)
	require.Equal(t, transaction1.PricePerCoin, transaction2.PricePerCoin)
	require.WithinDuration(t, transaction1.TimeTransacted, transaction2.TimeTransacted, time.Second)
	require.WithinDuration(t, transaction1.TimeCreated, transaction2.TimeCreated, time.Second)
}

func TestDeleteTransaction(t *testing.T) {
	user, transaction1 := createRandomTransaction(t)
	arg := DeleteTransactionParams{
		ID:        transaction1.ID,
		AccountID: user.ID,
	}
	err := testStore.DeleteTransaction(context.Background(), arg)
	require.NoError(t, err)

	transaction2, err := testStore.GetTransaction(context.Background(), GetTransactionParams(arg))
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, transaction2)
}

func createTransactionsForAccount(t *testing.T, user *User, portfolio *Portfolio) Transaction {
	coin := util.RandomString(3)

	arg := CreateTransactionParams{
		PortfolioID:    portfolio.ID,
		AccountID:      user.ID,
		Symbol:         coin,
		Type:           0,
		Quantity:       float64(util.RandomInt()),
		PricePerCoin:   float64(util.RandomInt()),
		TimeTransacted: time.Time{},
	}

	transaction, err := testStore.CreateTransaction(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transaction)

	require.Equal(t, arg.AccountID, transaction.AccountID)
	require.Equal(t, arg.PortfolioID, transaction.PortfolioID)
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
	userTransactions, err := testStore.ListTransactionsByAccount(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, userTransactions, 10)

	for _, userTransaction := range userTransactions {
		require.NotEmpty(t, userTransaction)
	}

}

func createTransactionsForAccountForCoin(t *testing.T, user *User, portfolio *Portfolio, coin string) Transaction {

	arg := CreateTransactionParams{
		AccountID:      user.ID,
		PortfolioID:    portfolio.ID,
		Symbol:         coin,
		Type:           1,
		Quantity:       float64(util.RandomInt()),
		PricePerCoin:   float64(util.RandomInt()),
		TimeTransacted: time.Time{},
	}

	transaction, err := testStore.CreateTransaction(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transaction)

	require.Equal(t, arg.AccountID, transaction.AccountID)
	require.Equal(t, arg.PortfolioID, transaction.PortfolioID)
	require.Equal(t, arg.Symbol, transaction.Symbol)
	require.Equal(t, arg.Type, transaction.Type)
	require.Equal(t, arg.Quantity, transaction.Quantity)
	require.Equal(t, arg.PricePerCoin, transaction.PricePerCoin)
	require.WithinDuration(t, arg.TimeTransacted, transaction.TimeTransacted, time.Second)

	return transaction
}

func TestListTransactionsByAccountByCoin(t *testing.T) {
	user, portfolio := createUserAndPortfolio(t)
	coin := util.RandomString(3)

	arg := ListTransactionsByAccountByCoinParams{
		AccountID: user.ID,
		Symbol:    coin,
		Limit:     10,
		Offset:    0,
	}

	for i := 0; i < 10; i++ {
		createTransactionsForAccountForCoin(t, &user, &portfolio, coin)
	}

	userTransactions, err := testStore.ListTransactionsByAccountByCoin(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, userTransactions, 10)

	for _, userTransaction := range userTransactions {
		require.NotEmpty(t, userTransaction)
	}
}

func TestGetRollUpByCoinByPortfolio(t *testing.T) {
	user, portfolio := createUserAndPortfolio(t)

	for i := 0; i < 10; i++ {
		coin := util.RandomString(3)
		createTransactionsForAccountForCoin(t, &user, &portfolio, coin)
	}
	arg := GetRollUpByCoinByPortfolioParams{
		PortfolioID: portfolio.ID,
		AccountID:   user.ID,
	}
	rollUp, err := testStore.GetRollUpByCoinByPortfolio(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, rollUp, 10)

	for _, rollUp := range rollUp {
		require.NotEmpty(t, rollUp)
	}
}
