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

func (server *Server) getCoin(ctx *gin.Context) {
	var req getCoinRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	coin, err := server.store.GetCoin(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, coin)
}

type updateCoinRequest struct {
	ID          int64  `json:"id" binding:"required,min=1"`
	Amount      int32  `json:"amount" binding:"required"`
	NoOfCoins   string `json:"no_of_coins" binding:"required"`
}

func (server *Server) updateCoin(ctx *gin.Context) {
	var req updateCoinRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateCoinParams{
		ID:          req.ID,
		Amount:      req.Amount,
		NoOfCoins:   req.NoOfCoins,
	}
	coin, err := server.store.UpdateCoin(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, coin)
}

type deleteCoinRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteCoin(ctx *gin.Context) {
	var req deleteCoinRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteCoin(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

type listCoinsRequest struct {
	PortfolioID int64 `uri:"portfolio_id" binding:"required,min=1"`
}

func (server *Server) listCoins(ctx *gin.Context) {
	var req listCoinsRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	coins, err := server.store.ListCoins(ctx, req.PortfolioID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, coins)
}

