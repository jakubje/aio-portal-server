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

func (server *Server) ListCoins(ctx context.Context, req *pb.ListCoinsRequest) (*pb.ListCoinsResponse, error) {
	//authPayload, err := server.authorizeUser(ctx)
	//if err != nil {
	//	return nil, unauthenticatedError(err)
	//}
	violations := validateListCoinsRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	arg := db.ListCoinsParams{
		Limit:  req.GetLimit(),
		Offset: req.GetOffset(),
	}
	coins, err := server.store.ListCoins(ctx, arg)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "coins not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to get coins: %s", err)
	}

	rsp := &pb.ListCoinsResponse{
		Coins: convertCoins(coins),
	}
	return rsp, nil
}

func validateListCoinsRequest(req *pb.ListCoinsRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateNumberLimit(req.GetLimit(), 0, 200); err != nil {
		violations = append(violations, fieldViolation("limit", err))
	}
	if err := val.ValidateNumberLimit(req.GetOffset(), 0, 200); err != nil {
		violations = append(violations, fieldViolation("offset", err))
	}

	return violations
}
