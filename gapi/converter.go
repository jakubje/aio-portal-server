package gapi

import (
	db "github.com/jakub/aioportal/server/db/sqlc"
	"github.com/jakub/aioportal/server/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	BOUGHT = int32(0)
	SOLD   = int32(1)
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

type portfolioRollUp struct {
	totalCost            float64
	totalCoins           float64
	pricePerCoin         float64
	profitLossPercentage float64
	currentValue         float64
}

func convertRollUp(rollUps []db.GetRollUpByCoinByPortfolioRow) []*pb.Roll_Up {
	var portfolio = make(map[string]portfolioRollUp)

	var rollUpProtoArray []*pb.Roll_Up
	//for _, rollUp := range rollUps {
	//	portfolio[rollUp.Symbol] = portfolioRollUp{
	//		totalCost:    0,
	//		totalCoins:   0,
	//		pricePerCoin: 0,
	//	}
	//}

	for _, rollUp := range rollUps {
		currentRollUp := portfolio[rollUp.Symbol]

		if rollUp.Type == BOUGHT {
			currentRollUp.totalCost += rollUp.Amount
			currentRollUp.totalCoins += rollUp.TotalCoins
		} else if rollUp.Type == SOLD {
			currentRollUp.totalCost -= rollUp.Amount
			currentRollUp.totalCoins -= rollUp.TotalCoins
		}
		currentRollUp.profitLossPercentage += rollUp.ProfitLossPercentage
		currentRollUp.currentValue = currentRollUp.totalCost * (1 + (currentRollUp.profitLossPercentage / 100))

		currentRollUp.pricePerCoin = currentRollUp.totalCost / currentRollUp.totalCoins
		portfolio[rollUp.Symbol] = currentRollUp
	}

	for key, coin := range portfolio {

		rollUpProtoArray = append(rollUpProtoArray, &pb.Roll_Up{
			Symbol:               key,
			TotalCost:            coin.totalCost,
			TotalCoins:           coin.totalCoins,
			PricePerCoin:         coin.pricePerCoin,
			ProfitLossPercentage: coin.profitLossPercentage,
			CurrentValue:         coin.currentValue,
		})
	}

	//for _, rollUp := range rollUps {
	//
	//	rollUpProtoArray = append(rollUpProtoArray, &pb.Roll_Up{
	//		Symbol:       rollUp.Symbol,
	//		TotalCost:    rollUp.TotalCost,
	//		Type:         rollUp.Type,
	//		TotalCoins:   rollUp.TotalCoins,
	//		PricePerCoin: rollUp.PricePerCoin,
	//	})
	//}

	return rollUpProtoArray
}

func convertTransaction(transaction db.Transaction) *pb.Transaction {
	transactionId := transaction.ID.String()
	return &pb.Transaction{
		Id:           transactionId,
		PortfolioId:  transaction.PortfolioID,
		Symbol:       transaction.Symbol,
		Type:         transaction.Type,
		Amount:       transaction.Amount,
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
			Amount:       transaction.Amount,
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

//type Tags struct {
//	Layer string
//	Pow   string
//}
//
//func converTags(tag Tags) *pb.Tags {
//	return &pb.Tags{
//		Layer: tag.Layer,
//		Pow:   tag.Pow,
//	}
//}
//
//func nullStringToTags(ns sql.NullString) (Tags, error) {
//	if !ns.Valid {
//		return Tags{}, nil
//	}
//
//	var tags Tags
//	err := json.Unmarshal([]byte(ns.String), &tags)
//	if err != nil {
//		return Tags{}, err
//	}
//
//	return tags, nil
//}

func convertCoin(coin db.Coin) *pb.Coin {

	return &pb.Coin{
		CoinId:            coin.CoinID,
		CoinUuid:          coin.CoinUuid,
		Name:              coin.Name,
		Price:             coin.Price,
		MarketCap:         coin.MarketCap,
		NumberOfMarkets:   coin.NumberOfMarkets,
		NumberOfExchanges: coin.NumberOfExchanges,
		ApprovedSupply:    coin.ApprovedSupply,
		CirculatingSupply: coin.CirculatingSupply,
		TotalSupply:       coin.TotalSupply,
		MaxSupply:         coin.MaxSupply,
		Rank:              coin.Rank,
		Volume:            coin.Volume,
		DailyChange:       coin.DailyChange,
		ImageUrl:          coin.ImageUrl,
		Description:       coin.Description,
		AllTimeHigh:       coin.AllTimeHigh,
		Tags:              coin.Tags,
		Website:           coin.Website,
		SocialMediaLinks:  coin.SocialMediaLinks,
		Updated_At:        timestamppb.New(coin.UpdatedAt),
		CreatedAt:         timestamppb.New(coin.CreatedAt),
	}
}

func convertCoins(coins []db.Coin) []*pb.Coin {
	var coinsProtoArray []*pb.Coin
	for _, coin := range coins {

		coinsProtoArray = append(coinsProtoArray, &pb.Coin{
			CoinId:            coin.CoinID,
			CoinUuid:          coin.CoinUuid,
			Name:              coin.Name,
			Price:             coin.Price,
			MarketCap:         coin.MarketCap,
			NumberOfMarkets:   coin.NumberOfMarkets,
			NumberOfExchanges: coin.NumberOfExchanges,
			CirculatingSupply: coin.CirculatingSupply,
			TotalSupply:       coin.TotalSupply,
			MaxSupply:         coin.MaxSupply,
			Rank:              coin.Rank,
			Volume:            coin.Volume,
			DailyChange:       coin.DailyChange,
			ImageUrl:          coin.ImageUrl,
			Description:       coin.Description,
			AllTimeHigh:       coin.AllTimeHigh,
			Tags:              coin.Tags,
			Website:           coin.Website,
			SocialMediaLinks:  coin.SocialMediaLinks,
			Updated_At:        timestamppb.New(coin.UpdatedAt),
			CreatedAt:         timestamppb.New(coin.CreatedAt),
		})
	}
	return coinsProtoArray
}
