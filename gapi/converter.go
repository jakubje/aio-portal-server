package gapi

import (
	db "github.com/jakub/aioportal/server/db/sqlc"
	"github.com/jakub/aioportal/server/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user db.User) *pb.User {
	return &pb.User{
		Email:             user.Email,
		Name:              user.Name,
		LastName:          user.LastName,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
		CreatedAt:         timestamppb.New(user.CreatedAt),
	}
}

func convertPortfolio(portfolio db.Portfolio) *pb.Portfolio {
	return &pb.Portfolio{
		Id:         portfolio.ID,
		AccountId:  portfolio.AccountID,
		Name:       portfolio.Name,
		Holdings:   portfolio.Holdings,
		Change_24H: portfolio.Change24h,
		ProfitLoss: portfolio.ProfitLoss,
	}
}
func convertPortfolios(portfolios []db.Portfolio) []*pb.Portfolio {

	var portfoliosArr []*pb.Portfolio
	for _, portfolio := range portfolios {
		portfoliosArr = append(portfoliosArr, &pb.Portfolio{
			Id:         portfolio.ID,
			AccountId:  portfolio.AccountID,
			Name:       portfolio.Name,
			Holdings:   portfolio.Holdings,
			Change_24H: portfolio.Change24h,
			ProfitLoss: portfolio.ProfitLoss,
		})
	}

	return portfoliosArr
}

func convertRollUp(rollUps []db.GetRollUpByCoinByPortfolioRow) []*pb.Roll_Up {

	var rollUpArr []*pb.Roll_Up
	for _, rollUp := range rollUps {
		rollUpArr = append(rollUpArr, &pb.Roll_Up{
			Symbol:       rollUp.Symbol,
			Type:         rollUp.Type,
			TotalCost:    float32(rollUp.TotalCost),
			TotalCoins:   float32(rollUp.TotalCoins),
			PricePerCoin: float32(rollUp.PricePerCoin),
		})
	}
	return rollUpArr
}
