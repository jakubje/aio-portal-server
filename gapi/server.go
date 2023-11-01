package gapi

import (
	"fmt"
	db "github.com/jakub/aioportal/server/db/sqlc"
	"github.com/jakub/aioportal/server/pb"
	"github.com/jakub/aioportal/server/token"
	"github.com/jakub/aioportal/server/util"
	"github.com/jakub/aioportal/server/worker"
)

type Server struct {
	pb.UnimplementedAioPortalServer
	config          util.Config
	store           db.Store
	tokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
}

// NewServer creates a new gRPC server
func NewServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	//tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:          config,
		store:           store,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
	}

	return server, nil
}
