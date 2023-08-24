package db

import (
	"context"

	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jakub/aioportal/server/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.RandomString(8))
	require.NoError(t, err)

	arg := CreateUserParams{
		Email:    util.RandomEmail(),
		Name:     util.RandomString(5),
		LastName: util.RandomString(5),
		Password: hashedPassword,
	}
	user, err := testStore.CreateUser(context.Background(), arg)
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
	account2, err := testStore.GetUser(context.Background(), account1.Email)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Email, account2.Email)
	require.Equal(t, account1.Name, account2.Name)
	require.Equal(t, account1.LastName, account2.LastName)
	require.Equal(t, account1.Password, account2.Password)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestUpdateUserName(t *testing.T) {
	account1 := createRandomUser(t)

	arg := UpdateUserParams{
		ID: account1.ID,
		Name: pgtype.Text{
			String: util.RandomString(5),
			Valid:  true,
		},
	}

	account2, err := testStore.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, arg.ID, account2.ID)
	require.NotEqual(t, account1.Name, account2.Name)
	require.Equal(t, account1.LastName, account2.LastName)
	require.Equal(t, account1.Password, account2.Password)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)

}

func TestUpdateUserEmail(t *testing.T) {
	account1 := createRandomUser(t)

	arg := UpdateUserParams{
		ID: account1.ID,
		Email: pgtype.Text{
			String: util.RandomEmail(),
			Valid:  true,
		},
	}

	account2, err := testStore.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, arg.ID, account2.ID)
	require.NotEqual(t, account1.Email, account2.Email)
	require.Equal(t, account1.Name, account2.Name)
	require.Equal(t, account1.LastName, account2.LastName)
	require.Equal(t, account1.Password, account2.Password)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)

}

func TestUpdateUserLastName(t *testing.T) {
	account1 := createRandomUser(t)

	arg := UpdateUserParams{
		ID: account1.ID,
		LastName: pgtype.Text{
			String: util.RandomString(7),
			Valid:  true,
		},
	}

	account2, err := testStore.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, arg.ID, account2.ID)
	require.NotEqual(t, account1.LastName, account2.LastName)
	require.Equal(t, account1.Email, account2.Email)
	require.Equal(t, account1.Name, account2.Name)
	require.Equal(t, account1.Password, account2.Password)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)

}

func TestUpdateUserPassword(t *testing.T) {
	account1 := createRandomUser(t)

	hashedPassword, err := util.HashPasswod(util.RandomString(10))
	require.NoError(t, err)

	arg := UpdateUserParams{
		ID: account1.ID,
		Password: pgtype.Text{
			String: hashedPassword,
			Valid:  true,
		},
	}

	account2, err := testStore.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, arg.ID, account2.ID)
	require.NotEqual(t, account1.Password, account2.Password)
	require.Equal(t, account1.Email, account2.Email)
	require.Equal(t, account1.Name, account2.Name)
	require.Equal(t, account1.LastName, account2.LastName)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)

}

func TestDeleteUser(t *testing.T) {
	account1 := createRandomUser(t)
	err := testStore.DeleteUser(context.Background(), account1.ID)
	require.NoError(t, err)

	account2, err := testStore.GetUser(context.Background(), account1.Email)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, account2)
}

func TestListUsers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomUser(t)
	}
	arg := ListUsersParams{
		Limit:  5,
		Offset: 0,
	}

	accounts, err := testStore.ListUsers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
