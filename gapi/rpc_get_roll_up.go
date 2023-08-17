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

func (server *Server) GetRollUp(ctx context.Context, req *pb.RollUpRequest) (*pb.RollUpResponse, error) {
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}
	violations := validateRollUpRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	arg := db.GetRollUpByCoinByPortfolioParams{
		PortfolioID: req.GetId(),
		AccountID:   authPayload.AccountId,
	}
	rollUp, err := server.store.GetRollUpByCoinByPortfolio(ctx, arg)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "no roll up found")
		}
		return nil, status.Errorf(codes.Internal, "failed to get roll up: %s", err)
	}

	rsp := &pb.RollUpResponse{
		Summary: convertRollUp(rollUp),
	}

	return rsp, nil
}

func validateRollUpRequest(req *pb.RollUpRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateID(req.GetId()); err != nil {
		violations = append(violations, fieldViolation("portfolio_id", err))
	}
	return violations
}
