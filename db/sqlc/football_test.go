package db

import (
	"context"
	"server/internal/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomFootball(t *testing.T) (Football, User) {
	user := createRandomUser(t)
	arg := CreateFootballParams{
		AccountID: user.ID,
		Team:      utils.RandomString(5),
		League:    utils.RandomString(5),
		Country:   utils.RandomString(5),
	}

	football, err := testQueries.CreateFootball(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, football)

	require.Equal(t, arg.AccountID, football.AccountID)
	require.Equal(t, arg.Team, football.Team)
	require.Equal(t, arg.League, football.League)
	require.Equal(t, arg.Country, football.Country)
	require.NotZero(t, football.ID)

	return football, user
}

func TestCreateFootball(t *testing.T) {
	createRandomFootball(t)
}

func TestGetFootball(t *testing.T) {

	football1, user := createRandomFootball(t)
	football2, err := testQueries.GetFootball(context.Background(), user.ID)
	require.NoError(t, err)
	require.NotEmpty(t, football2)

	require.Equal(t, football1.AccountID, football2.AccountID)
	require.Equal(t, football1.Team, football2.Team)
	require.Equal(t, football1.League, football2.League)
	require.Equal(t, football1.Country, football2.Country)

}

func TestUpdateFootball(t *testing.T) {
	_, user := createRandomFootball(t)

	arg := UpdateFootballParams{
		AccountID: user.ID,
		Team:      utils.RandomString(5),
		League:    utils.RandomString(5),
		Country:   utils.RandomString(5),
	}

	football2, err := testQueries.UpdateFootball(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, football2)

	require.Equal(t, arg.AccountID, football2.AccountID)
	require.Equal(t, arg.Team, football2.Team)
	require.Equal(t, arg.League, football2.League)
	require.Equal(t, arg.Country, football2.Country)
}
