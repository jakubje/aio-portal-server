package gapi

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jakub/aioportal/server/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (server *Server) ListPortfolios(ctx context.Context, _ *emptypb.Empty) (*pb.ListPortfoliosResponse, error) {
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	portfolios, err := server.store.ListPortforlios(ctx, authPayload.AccountId)
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
