package token

import (
	"github.com/jakub/aioportal/server/internal/utils"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestNewPasetoMaker(t *testing.T) {
	maker, err := NewPasetoMaker(utils.RandomString(32))
	require.NoError(t, err)

	userId := utils.RandomInt64()
	duration := time.Minute

	issuedAt := time.Now()
	expriedAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(userId, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, userId, payload.AccountId)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expriedAt, payload.ExpiredAt, time.Second)
}

func TestExpiredPasetoToken(t *testing.T) {
	maker, err := NewPasetoMaker(utils.RandomString(32))
	require.NoError(t, err)
	userId := utils.RandomInt64()
	token, err := maker.CreateToken(userId, -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)
}
