package db

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"server/internal/utils"
	"testing"
	"time"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Email:    utils.RandomEmail(),
		Name:     utils.RandomString(5),
		LastName: utils.RandomString(5),
		Password: utils.RandomString(8),
	}
	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Name, user.Name)
	require.Equal(t, arg.LastName, user.LastName)
	require.Equal(t, arg.Password, user.Password)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	//Create account
	account1 := createRandomUser(t)
	account2, err := testQueries.GetUser(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Email, account2.Email)
	require.Equal(t, account1.Name, account2.Name)
	require.Equal(t, account1.LastName, account2.LastName)
	require.Equal(t, account1.Password, account2.Password)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestUpdateUser(t *testing.T) {
	account1 := createRandomUser(t)

	arg := UpdateUserParams{
		ID:       account1.ID,
		Name:     utils.RandomString(5),
		LastName: utils.RandomString(5),
		Password: utils.RandomString(8),
	}

	account2, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, arg.ID, account2.ID)
	require.Equal(t, arg.Name, account2.Name)
	require.Equal(t, arg.LastName, account2.LastName)
	require.Equal(t, arg.Password, account2.Password)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)

}

func TestDeleteUser(t *testing.T) {
	account1 := createRandomUser(t)
	err := testQueries.DeleteUser(context.Background(), account1.ID)
	require.NoError(t, err)

	account2, err := testQueries.GetUser(context.Background(), account1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}

func TestListUsers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomUser(t)
	}
	arg := ListUsersParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListUsers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
