package gapi

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"time"

	db "github.com/jakub/aioportal/server/db/sqlc"
	"github.com/jakub/aioportal/server/pb"
	"github.com/jakub/aioportal/server/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) UpdateCoin(ctx context.Context, req *pb.UpdateCoinRequest) (*pb.UpdateCoinResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	violations := validateUpdateCoinRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	arg := db.UpdateCoinParams{
		Name: pgtype.Text{
			String: req.GetName(),
			Valid:  true,
		},
		Price: pgtype.Text{
			String: req.GetPrice(),
			Valid:  true,
		},
		MarketCap: pgtype.Text{
			String: req.GetMarketCap(),
			Valid:  true,
		},
		CirculatingSupply: pgtype.Text{
			String: req.GetCirculatingSupply(),
			Valid:  true,
		},
		TotalSupply: pgtype.Text{
			String: req.GetTotalSupply(),
			Valid:  true,
		},
		MaxSupply: pgtype.Text{
			String: req.GetMaxSupply(),
			Valid:  true,
		},
		Rank: pgtype.Text{
			String: req.GetRank(),
			Valid:  true,
		},
		Volume: pgtype.Text{
			String: req.GetVolume(),
			Valid:  true,
		},
		ImageUrl: pgtype.Text{
			String: req.GetImageUrl(),
			Valid:  true,
		},
		Description: pgtype.Text{
			String: req.GetDescription(),
			Valid:  true,
		},
		Website: pgtype.Text{
			String: req.GetWebsite(),
			Valid:  true,
		},
		SocialMediaLinks: req.GetSocialMediaLinks(),
		CoinID:           req.GetCoinId(),
		UpdatedAt: pgtype.Timestamptz{
			Time:  time.Now(),
			Valid: true,
		},
	}

	updatedCoin, err := server.store.UpdateCoin(ctx, arg)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "coin not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to update coin: %s", err)
	}

	rsp := &pb.UpdateCoinResponse{
		Coin: convertCoin(updatedCoin),
	}
	return rsp, nil
}

func validateUpdateCoinRequest(req *pb.UpdateCoinRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateString(req.GetCoinId(), 1, 20); err != nil {
		violations = append(violations, fieldViolation("coin_id", err))
	}
	return violations
}
