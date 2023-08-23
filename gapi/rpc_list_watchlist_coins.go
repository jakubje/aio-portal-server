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

func (server *Server) ListWatchlistCoins(ctx context.Context, req *pb.ListWatchlistCoinsRequest) (*pb.ListWatchlistCoinsResponse, error) {
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	violations := validateListWatchlistCoinsequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	arg := db.ListWatchlistCoinsParams{
		WatchlistID: req.GetWatchlistId(),
		AccountID:   authPayload.AccountId,
	}

	watchlistCoins, err := server.store.ListWatchlistCoins(ctx, arg)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "no coins found")
		}
		return nil, status.Errorf(codes.Internal, "failed to get watchlist coins: %s", err)
	}

	rsp := &pb.ListWatchlistCoinsResponse{
		Total: int64(len(watchlistCoins)),
		Coins: convertCoins(watchlistCoins),
	}
	return rsp, nil
}

func validateListWatchlistCoinsequest(req *pb.ListWatchlistCoinsRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateID(req.GetWatchlistId()); err != nil {
		violations = append(violations, fieldViolation("watchlist_id", err))
	}
	return violations
}
