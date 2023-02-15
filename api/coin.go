package api

import (
	"net/http"
	db "server/db/sqlc"

	"github.com/gin-gonic/gin"
)


type addCoinRequest struct {
	PortfolioID int64 `json:"portfolio_id" binding:"required"`
	CoinName   string `json:"coin_name" binding:"required"`
	CoinSymbol string `json:"coin_symbol" binding:"required"`
	Amount     int32 `json:"amount" binding:"required"`
	NoOfCoins  string `json:"no_of_coins" binding:"required"`
}


func (server *Server) addCoin(ctx *gin.Context) {
	var req addCoinRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.AddCoinParams{
		PortfolioID: req.PortfolioID,
		CoinName:    req.CoinName,
		CoinSymbol:  req.CoinSymbol,
		Amount:      req.Amount,
		NoOfCoins:   req.NoOfCoins,
	}
	coin, err := server.store.AddCoin(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, coin)
}

type getCoinRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

