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

func (server *Server) UpdatePortfolio(ctx context.Context, req *pb.UpdatePortfolioRequest) (*pb.UpdatePortfolioResponse, error) {
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	violations := validateUpdatePortfolioRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	arg := db.UpdatePortfolioParams{
		ID:        req.GetId(),
		Name:      req.GetName(),
		AccountID: authPayload.AccountId,
	}

	updatedPortfolio, err := server.store.UpdatePortfolio(ctx, arg)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "portfolio not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to get portfolio: %s", err)
	}

	rsp := &pb.UpdatePortfolioResponse{
		Portfolio: convertPortfolio(updatedPortfolio),
	}
	return rsp, nil
}

func validateUpdatePortfolioRequest(req *pb.UpdatePortfolioRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateString(req.GetName(), 1, 20); err != nil {
		violations = append(violations, fieldViolation("name", err))
	}
	return violations
}
