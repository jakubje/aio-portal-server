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

	var portfoliosProtoArray []*pb.Portfolio
	for _, portfolio := range portfolios {
		portfoliosProtoArray = append(portfoliosProtoArray, &pb.Portfolio{
			Id:         portfolio.ID,
			AccountId:  portfolio.AccountID,
			Name:       portfolio.Name,
			Holdings:   portfolio.Holdings,
			Change_24H: portfolio.Change24h,
			ProfitLoss: portfolio.ProfitLoss,
		})
	}

	return portfoliosProtoArray
}

func convertRollUp(rollUps []db.GetRollUpByCoinByPortfolioRow) []*pb.Roll_Up {

	var rollUpProtoArray []*pb.Roll_Up
	for _, rollUp := range rollUps {
		rollUpProtoArray = append(rollUpProtoArray, &pb.Roll_Up{
			Symbol:       rollUp.Symbol,
			Type:         rollUp.Type,
			TotalCost:    rollUp.TotalCost,
			TotalCoins:   rollUp.TotalCoins,
			PricePerCoin: rollUp.PricePerCoin,
		})
	}
	return rollUpProtoArray
}

func convertTransaction(transaction db.Transaction) *pb.Transaction {
	transactionId := transaction.ID.String()
	return &pb.Transaction{
		Id:           transactionId,
		PortfolioId:  transaction.PortfolioID,
		Symbol:       transaction.Symbol,
		Type:         transaction.Type,
		PricePerCoin: transaction.PricePerCoin,
		Quantity:     transaction.Quantity,
	}
}

func convertTransactions(transactions []db.Transaction) []*pb.Transaction {

	var transactionsProtoArray []*pb.Transaction
	for _, transaction := range transactions {
		transactionId := transaction.ID.String()
		transactionsProtoArray = append(transactionsProtoArray, &pb.Transaction{
			Id:           transactionId,
			PortfolioId:  transaction.PortfolioID,
			Symbol:       transaction.Symbol,
			Type:         transaction.Type,
			PricePerCoin: transaction.PricePerCoin,
			Quantity:     transaction.Quantity,
		})
	}

	return transactionsProtoArray
}

func convertWatchlist(watchlist db.Watchlist) *pb.Watchlist {
	return &pb.Watchlist{
		Id:   watchlist.ID,
		Name: watchlist.Name,
	}
}

func convertWatchlists(watchlists []db.Watchlist) []*pb.Watchlist {

	var watchlistsProtoArray []*pb.Watchlist
	for _, watchlist := range watchlists {
		watchlistsProtoArray = append(watchlistsProtoArray, &pb.Watchlist{
			Id:   watchlist.ID,
			Name: watchlist.Name,
		})
	}
	return watchlistsProtoArray
}
