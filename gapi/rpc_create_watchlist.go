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

func (server *Server) CreateWatchlist(ctx context.Context, req *pb.CreateWatchlistRequest) (*pb.CreateWatchlistResponse, error) {
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	violations := validateCreateWatchlistRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	arg := db.CreateWatchlistParams{
		Name:      req.GetName(),
		AccountID: authPayload.AccountId,
	}

	watchlist, err := server.store.CreateWatchlist(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create watchlist: %s", err)
	}

	rsp := &pb.CreateWatchlistResponse{
		Watchlist: convertWatchlist(watchlist),
	}
	return rsp, nil
}

func validateCreateWatchlistRequest(req *pb.CreateWatchlistRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateString(req.GetName(), 1, 20); err != nil {
		violations = append(violations, fieldViolation("name", err))
	}
	return violations
}
