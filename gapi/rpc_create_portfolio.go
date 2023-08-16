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

func (server *Server) CreatePortfolio(ctx context.Context, req *pb.CreatePortfolioRequest) (*pb.CreatePortfolioResponse, error) {
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	violations := validateCreatePortfolioRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	arg := db.CreatePortfolioParams{
		Name:      req.GetName(),
		AccountID: authPayload.AccountId,
	}

	portfolio, err := server.store.CreatePortfolio(ctx, arg)
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			return nil, status.Errorf(codes.AlreadyExists, "email already exists: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "failed to create user: %s", err)
	}

	rsp := &pb.CreatePortfolioResponse{
		Portfolio: convertPortfolio(portfolio),
	}
	return rsp, nil
}

func validateCreatePortfolioRequest(req *pb.CreatePortfolioRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateString(req.GetName(), 1, 20); err != nil {
		violations = append(violations, fieldViolation("name", err))
	}
	return violations
}
