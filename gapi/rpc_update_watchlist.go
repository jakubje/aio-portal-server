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

func (server *Server) UpdateWatchlist(ctx context.Context, req *pb.UpdateWatchlistRequest) (*pb.UpdateWatchlistResponse, error) {
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	violations := validateUpdateWatchlistRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	arg := db.UpdateWatchlistParams{
		ID:        req.GetId(),
		Name:      req.GetName(),
		AccountID: authPayload.AccountId,
	}

	updatedWatchlist, err := server.store.UpdateWatchlist(ctx, arg)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "Watchlist not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to get Watchlist: %s", err)
	}

	rsp := &pb.UpdateWatchlistResponse{
		Watchlist: convertWatchlist(updatedWatchlist),
	}
	return rsp, nil
}

func validateUpdateWatchlistRequest(req *pb.UpdateWatchlistRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateString(req.GetName(), 1, 20); err != nil {
		violations = append(violations, fieldViolation("name", err))
	}
	return violations
}
