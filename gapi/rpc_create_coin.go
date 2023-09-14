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

func (server *Server) CreateCoin(ctx context.Context, req *pb.CreateCoinRequest) (*pb.CreateCoinResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	violations := validateAddCoinRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	arg := db.CreateCoinParams{
		CoinID:            req.GetCoinId(),
		CoinUuid:          req.GetCoinUuid(),
		Name:              req.GetName(),
		Price:             req.GetPrice(),
		MarketCap:         req.GetMarketCap(),
		NumberOfMarkets:   req.GetNumberOfMarkets(),
		NumberOfExchanges: req.GetNumberOfExchanges(),
		ApprovedSupply:    req.GetApprovedSupply(),
		CirculatingSupply: req.GetCirculatingSupply(),
		TotalSupply:       req.GetTotalSupply(),
		MaxSupply:         req.GetMaxSupply(),
		Rank:              req.GetRank(),
		Volume:            req.GetVolume(),
		DailyChange:       req.GetDailyChange(),
		ImageUrl:          req.GetImageUrl(),
		Description:       req.GetDescription(),
		AllTimeHigh:       req.GetAllTimeHigh(),
		Tags:              req.GetTags(),
		Website:           req.GetWebsite(),
		SocialMediaLinks:  req.GetSocialMediaLinks(),
	}

	coin, err := server.store.CreateCoin(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create coin: %s", err)
	}

	rsp := &pb.CreateCoinResponse{
		Coin: convertCoin(coin),
	}
	return rsp, nil
}

func validateAddCoinRequest(req *pb.CreateCoinRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateString(req.GetCoinId(), 1, 20); err != nil {
		violations = append(violations, fieldViolation("coin_id", err))
	}
	//if err := val.ValidateString(req.GetName(), 1, 20); err != nil {
	//	violations = append(violations, fieldViolation("name", err))
	//}
	//if err := val.ValidateFloat(req.GetPrice()); err != nil {
	//	violations = append(violations, fieldViolation("price", err))
	//}
	//if err := val.ValidateInt(req.GetMarketCap()); err != nil {
	//	violations = append(violations, fieldViolation("market_cap", err))
	//}
	//if err := val.ValidateInt(req.GetCirculatingSupply()); err != nil {
	//	violations = append(violations, fieldViolation("circulating_supply", err))
	//}
	//if err := val.ValidateInt(req.GetTotalSupply()); err != nil {
	//	violations = append(violations, fieldViolation("total_supply", err))
	//}
	//if err := val.ValidateInt(req.GetMaxSupply()); err != nil {
	//	violations = append(violations, fieldViolation("max_supply", err))
	//}
	//if err := val.ValidateInt(req.GetRank()); err != nil {
	//	violations = append(violations, fieldViolation("rank", err))
	//}
	//if err := val.ValidateInt(req.GetVolume()); err != nil {
	//	violations = append(violations, fieldViolation("volume", err))
	//}
	//if err := val.ValidateString(req.GetImageUrl(), 1, 20); err != nil {
	//	violations = append(violations, fieldViolation("image_url", err))
	//}
	//if err := val.ValidateString(req.GetDescription(), 1, 20); err != nil {
	//	violations = append(violations, fieldViolation("description", err))
	//}
	//if err := val.ValidateString(req.GetWebsite(), 1, 20); err != nil {
	//	violations = append(violations, fieldViolation("website", err))
	//}
	//if err := val.ValidateSocialMediaLinks(req.GetSocialMediaLinks()); err != nil {
	//	violations = append(violations, fieldViolation("social_media_links", err))
	//}
	return violations
}
