package gapi

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jakub/aioportal/server/pb"
	"github.com/jakub/aioportal/server/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetCoin(ctx context.Context, req *pb.GetCoinRequest) (*pb.GetCoinResponse, error) {

	violations := validateCoinRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	coin, err := server.store.GetCoin(ctx, req.GetCoinId())
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "coin not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to get coin: %s", err)
	}

	rsp := &pb.GetCoinResponse{
		Coin: convertCoin(coin),
	}
	return rsp, nil
}

func validateCoinRequest(req *pb.GetCoinRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateString(req.GetCoinId(), 2, 6); err != nil {
		violations = append(violations, fieldViolation("coin_id", err))
	}

	return violations
}
