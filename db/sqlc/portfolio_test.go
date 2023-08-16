package db

import (
	"context"

	"testing"

	"github.com/jakub/aioportal/server/util"
	"github.com/stretchr/testify/require"
)

func createRandomPortfolio(t *testing.T) Portfolio {
	user := createRandomUser(t)
	arg := CreatePortfolioParams{
		Name:      util.RandomString(5),
		AccountID: user.ID,
	}
	portfolio, err := testStore.CreatePortfolio(context.Background(), arg)
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

func createUserAndPortfolio(t *testing.T) (User, Portfolio) {
	user := createRandomUser(t)
	arg := CreatePortfolioParams{
		Name:      util.RandomString(5),
		AccountID: user.ID,
	}
	portfolio, err := testStore.CreatePortfolio(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, portfolio)

	require.Equal(t, arg.Name, portfolio.Name)
	require.Equal(t, arg.AccountID, portfolio.AccountID)

	require.NotZero(t, portfolio.ID)

	return user, portfolio
}

func TestCreateUserAndPortfolio(t *testing.T) {
	createUserAndPortfolio(t)
}

func TestGetPortfolio(t *testing.T) {
	user, portfolio1 := createUserAndPortfolio(t)
	arg := GetPortfolioParams{
		ID:        portfolio1.ID,
		AccountID: user.ID,
	}
	portfolio2, err := testStore.GetPortfolio(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, portfolio2)

	require.Equal(t, portfolio1.ID, portfolio2.ID)
	require.Equal(t, portfolio1.AccountID, portfolio2.AccountID)
	require.Equal(t, portfolio1.Name, portfolio2.Name)
}

func TestUpdatePortfolio(t *testing.T) {
	user, portfolio1 := createUserAndPortfolio(t)

	arg := UpdatePortfolioParams{
		ID:        portfolio1.ID,
		Name:      util.RandomString(5),
		AccountID: user.ID,
	}

	portfolio2, err := testStore.UpdatePortfolio(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, portfolio2)

	require.Equal(t, arg.ID, portfolio2.ID)
	require.Equal(t, arg.Name, portfolio2.Name)

}

func TestDeletePortfolio(t *testing.T) {
	user, portfolio1 := createUserAndPortfolio(t)

	arg := DeletePortfolioParams{
		ID:        portfolio1.ID,
		AccountID: user.ID,
	}
	err := testStore.DeletePortfolio(context.Background(), arg)
	require.NoError(t, err)

	portfolio2, err := testStore.GetPortfolio(context.Background(), GetPortfolioParams(arg))
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, portfolio2)

}

func createPortfoliosByUser(t *testing.T, user *User) Portfolio {
	arg := CreatePortfolioParams{
		Name:      util.RandomString(5),
		AccountID: user.ID,
	}
	portfolio, err := testStore.CreatePortfolio(context.Background(), arg)
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

	userPortfolios, err := testStore.ListPortforlios(context.Background(), user.ID)
	require.NoError(t, err)
	require.Len(t, userPortfolios, 10)

	for _, userPortfolio := range userPortfolios {
		require.NotEmpty(t, userPortfolio)
	}
}
