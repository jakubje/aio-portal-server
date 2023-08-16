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

func (server *Server) ListPortfolios(ctx context.Context, req *pb.ListPortfoliosRequest) (*pb.ListPortfoliosResponse, error) {
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}
	req.Id = authPayload.AccountId
	violations := validateListPortfoliosRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	portfolios, err := server.store.ListPortforlios(ctx, req.GetId())
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "no portfolios found")
		}
		return nil, status.Errorf(codes.Internal, "failed to get portfolios: %s", err)
	}

	rsp := &pb.ListPortfoliosResponse{
		Portfolios: convertPortfolios(portfolios),
	}
	return rsp, nil
}

func validateListPortfoliosRequest(req *pb.ListPortfoliosRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateID(req.GetId()); err != nil {
		violations = append(violations, fieldViolation("portfolio_id", err))
	}

	return violations
}
