package db

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"server/internal/utils"
	"testing"
)

func createRandomPortfolio(t *testing.T) Portfolio {
	user := createRandomUser(t)
	arg := CreatePortfolioParams{
		Name:      utils.RandomString(5),
		AccountID: user.ID,
	}
	portfolio, err := testQueries.CreatePortfolio(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, portfolio)

	require.Equal(t, arg.Name, portfolio.Name)
	require.Equal(t, arg.AccountID, portfolio.AccountID)

	require.NotZero(t, portfolio.ID)

	return portfolio
}

func TestCreatePortfolio(t *testing.T) {
	createRandomPortfolio(t)
}

func TestGetPortfolio(t *testing.T) {
	portfolio1 := createRandomPortfolio(t)
	portfolio2, err := testQueries.GetPortfolio(context.Background(), portfolio1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, portfolio2)

	require.Equal(t, portfolio1.ID, portfolio2.ID)
	require.Equal(t, portfolio1.Name, portfolio2.Name)
	require.Equal(t, portfolio1.AccountID, portfolio2.AccountID)
}

func TestUpdatePortfolio(t *testing.T) {
	portfolio1 := createRandomPortfolio(t)

	arg := UpdatePortfolioParams{
		ID:   portfolio1.ID,
		Name: utils.RandomString(5),
	}

	portfolio2, err := testQueries.UpdatePortfolio(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, portfolio2)

	require.Equal(t, arg.ID, portfolio2.ID)
	require.Equal(t, arg.Name, portfolio2.Name)

}

func TestDeletePortfolio(t *testing.T) {
	portfolio1 := createRandomPortfolio(t)
	err := testQueries.DeletePortfolio(context.Background(), portfolio1.ID)
	require.NoError(t, err)

	portfolio2, err := testQueries.GetPortfolio(context.Background(), portfolio1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, portfolio2)

}

func createPortfoliosByUser(t *testing.T, user *User) Portfolio {
	arg := CreatePortfolioParams{
		Name:      utils.RandomString(5),
		AccountID: user.ID,
	}
	portfolio, err := testQueries.CreatePortfolio(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, portfolio)

	require.Equal(t, arg.Name, portfolio.Name)
	require.Equal(t, arg.AccountID, portfolio.AccountID)

	require.NotZero(t, portfolio.ID)

	return portfolio
}

func TestListPortfoliosByUser(t *testing.T) {
	user := createRandomUser(t)

	for i := 0; i < 10; i++ {
		createPortfoliosByUser(t, &user)
	}

	userPortfolios, err := testQueries.ListPortforlios(context.Background(), user.ID)
	require.NoError(t, err)
	require.Len(t, userPortfolios, 10)

	for _, userPortfolio := range userPortfolios {
		require.NotEmpty(t, userPortfolio)
	}
}
