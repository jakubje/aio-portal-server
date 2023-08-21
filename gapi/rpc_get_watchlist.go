package gapi

import (
	"context"
	"github.com/jackc/pgx/v5"
	db "github.com/jakub/aioportal/server/db/sqlc"
	"github.com/jakub/aioportal/server/pb"
	"github.com/jakub/aioportal/server/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetWatchlist(ctx context.Context, req *pb.GetWatchlistRequest) (*pb.GetWatchlistResponse, error) {
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}
	violations := validateWatchlistRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	arg := db.GetWatchlistParams{
		ID:        req.GetId(),
		AccountID: authPayload.AccountId,
	}

	watchlist, err := server.store.GetWatchlist(ctx, arg)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "Watchlist not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to get Watchlist: %s", err)
	}

	rsp := &pb.GetWatchlistResponse{
		Watchlist: convertWatchlist(watchlist),
	}
	return rsp, nil
}

func validateWatchlistRequest(req *pb.GetWatchlistRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateID(req.GetId()); err != nil {
		violations = append(violations, fieldViolation("Watchlist_id", err))
	}

	return violations
}
