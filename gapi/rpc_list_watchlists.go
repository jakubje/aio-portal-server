package gapi

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jakub/aioportal/server/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (server *Server) ListWatchlists(ctx context.Context, _ *emptypb.Empty) (*pb.GetWatchlistsResponse, error) {
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	Watchlists, err := server.store.ListWatchlists(ctx, authPayload.AccountId)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "Watchlists not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to get Watchlists: %s", err)
	}

	rsp := &pb.GetWatchlistsResponse{
		Watchlists: convertWatchlists(Watchlists),
	}
	return rsp, nil
}
