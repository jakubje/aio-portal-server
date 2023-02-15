package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/stretchr/testify/require"
	"server/internal/utils"
	"strconv"
	"testing"
	"time"
)

func createRandomTransaction(t *testing.T) Transaction {
	user := createRandomUser(t)
	coin := createRandomPortfolioCoin(t)
	fmt.Println(coin.ID)

	arg := CreateTransactionParams{
		AccountID:        user.ID,
		CoinID:           coin.ID,
		CoinName:         coin.CoinName,
		Symbol:           coin.CoinSymbol,
		Type:             0,
		Amount:           int32(utils.RandomInt()),
		TimeTransacted:   time.Time{},
		PricePurchasedAt: strconv.Itoa(utils.RandomInt()),
		NoOfCoins:        strconv.Itoa(utils.RandomInt()),
	}

	transaction, err := testQueries.CreateTransaction(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transaction)

	require.Equal(t, arg.AccountID, transaction.AccountID)
	require.Equal(t, arg.CoinID, transaction.CoinID)
	require.Equal(t, arg.CoinName, transaction.CoinName)
	require.Equal(t, arg.Symbol, transaction.Symbol)
	require.Equal(t, arg.Type, transaction.Type)
	require.Equal(t, arg.Amount, transaction.Amount)
	require.Equal(t, arg.PricePurchasedAt, transaction.PricePurchasedAt)
	require.Equal(t, arg.NoOfCoins, transaction.NoOfCoins)
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
	require.Equal(t, transaction1.CoinID, transaction2.CoinID)
	require.Equal(t, transaction1.CoinName, transaction2.CoinName)
	require.Equal(t, transaction1.Symbol, transaction2.Symbol)
	require.Equal(t, transaction1.Type, transaction2.Type)
	require.Equal(t, transaction1.Amount, transaction2.Amount)
	require.Equal(t, transaction1.PricePurchasedAt, transaction2.PricePurchasedAt)
	require.Equal(t, transaction1.NoOfCoins, transaction2.NoOfCoins)
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

func createTransactionsForAccount(t *testing.T, user *User) Transaction {
	coin := createRandomPortfolioCoin(t)

	arg := CreateTransactionParams{
		AccountID:        user.ID,
		CoinID:           coin.ID,
		CoinName:         coin.CoinName,
		Symbol:           coin.CoinSymbol,
		Type:             0,
		Amount:           int32(utils.RandomInt()),
		TimeTransacted:   time.Time{},
		PricePurchasedAt: strconv.Itoa(utils.RandomInt()),
		NoOfCoins:        strconv.Itoa(utils.RandomInt()),
	}

	transaction, err := testQueries.CreateTransaction(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transaction)

	require.Equal(t, arg.AccountID, transaction.AccountID)
	require.Equal(t, arg.CoinID, transaction.CoinID)
	require.Equal(t, arg.CoinName, transaction.CoinName)
	require.Equal(t, arg.Symbol, transaction.Symbol)
	require.Equal(t, arg.Type, transaction.Type)
	require.Equal(t, arg.Amount, transaction.Amount)
	require.Equal(t, arg.PricePurchasedAt, transaction.PricePurchasedAt)
	require.Equal(t, arg.NoOfCoins, transaction.NoOfCoins)
	require.WithinDuration(t, arg.TimeTransacted, transaction.TimeTransacted, time.Second)

	return transaction
}

func TestListTransactionsByAccount(t *testing.T) {
	user := createRandomUser(t)

	for i := 0; i < 10; i++ {
		createTransactionsForAccount(t, &user)
	}

	userTransactions, err := testQueries.ListTransactionsByAccount(context.Background(), user.ID)
	require.NoError(t, err)
	require.Len(t, userTransactions, 10)

	for _, userTransaction := range userTransactions {
		require.NotEmpty(t, userTransaction)
	}

}

func createTransactionsForAccountForCoin(t *testing.T, user *User, coin *Coin) Transaction {

	arg := CreateTransactionParams{
		AccountID:        user.ID,
		CoinID:           coin.ID,
		CoinName:         coin.CoinName,
		Symbol:           coin.CoinSymbol,
		Type:             0,
		Amount:           int32(utils.RandomInt()),
		TimeTransacted:   time.Time{},
		PricePurchasedAt: strconv.Itoa(utils.RandomInt()),
		NoOfCoins:        strconv.Itoa(utils.RandomInt()),
	}

	transaction, err := testQueries.CreateTransaction(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transaction)

	require.Equal(t, arg.AccountID, transaction.AccountID)
	require.Equal(t, arg.CoinID, transaction.CoinID)
	require.Equal(t, arg.CoinName, transaction.CoinName)
	require.Equal(t, arg.Symbol, transaction.Symbol)
	require.Equal(t, arg.Type, transaction.Type)
	require.Equal(t, arg.Amount, transaction.Amount)
	require.Equal(t, arg.PricePurchasedAt, transaction.PricePurchasedAt)
	require.Equal(t, arg.NoOfCoins, transaction.NoOfCoins)
	require.WithinDuration(t, arg.TimeTransacted, transaction.TimeTransacted, time.Second)

	return transaction
}

func TestListTransactionsByAccountByCoin(t *testing.T) {
	user := createRandomUser(t)
	coin := createRandomPortfolioCoin(t)

	arg := ListTransactionsByAccountByCoinParams{
		AccountID: user.ID,
		CoinID:    coin.ID,
	}

	for i := 0; i < 10; i++ {
		createTransactionsForAccountForCoin(t, &user, &coin)
	}

	userTransactions, err := testQueries.ListTransactionsByAccountByCoin(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, userTransactions, 10)

	for _, userTransaction := range userTransactions {
		require.NotEmpty(t, userTransaction)
	}
}
