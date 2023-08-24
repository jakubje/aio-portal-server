package gapi

import (
	"context"
	db "github.com/jakub/aioportal/server/db/sqlc"
	"github.com/jakub/aioportal/server/pb"
	"github.com/jakub/aioportal/server/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) AddWatchlistCoin(ctx context.Context, req *pb.AddWatchlistCoinRequest) (*pb.AddWatchlistCoinResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	violations := validateAddCointRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	arg := db.AddWatchlistCoinParams{
		WatchlistID: req.GetWatchlistId(),
		CoinID:      req.GetCoinId(),
	}

	_, err = server.store.AddWatchlistCoin(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create user: %s", err)
	}

	rsp := &pb.AddWatchlistCoinResponse{
		Message: "Coin has been added to your watchlist",
	}
	return rsp, nil
}

func validateAddCointRequest(req *pb.AddWatchlistCoinRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateString(req.GetCoinId(), 1, 20); err != nil {
		violations = append(violations, fieldViolation("coin_id", err))
	}
	if err := val.ValidateInt(req.GetWatchlistId()); err != nil {
		violations = append(violations, fieldViolation("watchlist_id", err))
	}
	return violations
}
