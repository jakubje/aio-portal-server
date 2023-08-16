package gapi

import (
	"context"
	"fmt"
	"testing"
	"time"

	db "github.com/jakub/aioportal/server/db/sqlc"
	"github.com/jakub/aioportal/server/token"
	"github.com/jakub/aioportal/server/util"
	"github.com/jakub/aioportal/server/worker"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"
)

func newTestServer(t *testing.T, store db.Store, taskDistributor worker.TaskDistributor) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store, taskDistributor)
	require.NoError(t, err)

	return server
}

func newContextWithBearerToken(t *testing.T, tokenMaker token.Maker, userId int64, userEmail string, duration time.Duration) context.Context {
	accessToken, _, err := tokenMaker.CreateToken(userId, userEmail, duration)
	require.NoError(t, err)
	bearerToken := fmt.Sprintf("%s %s", authorizationBearer, accessToken)
	md := metadata.MD{
		authorizationHeader: []string{
			bearerToken,
		},
	}
	return metadata.NewIncomingContext(context.Background(), md)
}
