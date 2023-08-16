package token

import (
	"testing"
	"time"

	"github.com/jakub/aioportal/server/util"
	"github.com/stretchr/testify/require"
)

func TestNewPasetoMaker(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	userId := util.RandomInt64(1, 1000)
	duration := time.Minute

	issuedAt := time.Now()
	expriedAt := issuedAt.Add(duration)
	userEmail := util.RandomEmail()

	token, payload, err := maker.CreateToken(userId, userEmail, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	payload, err = maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, userId, payload.AccountId)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expriedAt, payload.ExpiredAt, time.Second)
}

func TestExpiredPasetoToken(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)
	userId := util.RandomInt64(1, 1000)
	userEmail := util.RandomEmail()
	token, payload, err := maker.CreateToken(userId, userEmail, -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	payload, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)
}
