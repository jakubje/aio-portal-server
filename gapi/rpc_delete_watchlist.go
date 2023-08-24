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

func (server *Server) DeleteWatchlist(ctx context.Context, req *pb.DeleteWatchlistRequest) (*pb.DeleteWatchlistResponse, error) {
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}
	violations := validateWatchlistDeleteRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}
	arg := db.DeleteWatchlistParams{
		ID:        req.GetId(),
		AccountID: authPayload.AccountId,
	}
	err = server.store.DeleteWatchlist(ctx, arg)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "watchlist could not be deleted")
		}
		return nil, status.Errorf(codes.Internal, "failed to delete watchlist: %s", err)
	}

	rsp := &pb.DeleteWatchlistResponse{
		Message: "watchlist has been removed",
	}
	return rsp, nil
}

func validateWatchlistDeleteRequest(req *pb.DeleteWatchlistRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateID(req.GetId()); err != nil {
		violations = append(violations, fieldViolation("portfolio_id", err))
	}

	return violations
}
