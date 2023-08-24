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

func (server *Server) GetPortfolio(ctx context.Context, req *pb.GetPortfolioRequest) (*pb.GetPortfolioResponse, error) {
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}
	violations := validatePortfolioRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	arg := db.GetPortfolioParams{
		ID:        req.GetId(),
		AccountID: authPayload.AccountId,
	}
	portfolio, err := server.store.GetPortfolio(ctx, arg)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "portfolio not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to get portfolio: %s", err)
	}

	rsp := &pb.GetPortfolioResponse{
		Portfolio: convertPortfolio(portfolio),
	}
	return rsp, nil
}

func validatePortfolioRequest(req *pb.GetPortfolioRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateID(req.GetId()); err != nil {
		violations = append(violations, fieldViolation("portfolio_id", err))
	}

	return violations
}
