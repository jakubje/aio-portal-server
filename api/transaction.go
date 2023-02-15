package api

import (
	"net/http"
	db "server/db/sqlc"
	"time"

	"github.com/gin-gonic/gin"
)

type createTransactionRequest struct {
	AccountID        int64     `json:"account_id"`
	CoinID           int64     `json:"coin_id"`
	CoinName         string    `json:"coin_name"`
	Symbol           string    `json:"symbol"`
	Type             int32     `json:"type"`
	Amount           int32     `json:"amount"`
	TimeTransacted   time.Time `json:"time_transacted"`
	PricePurchasedAt string    `json:"price_purchased_at"`
	NoOfCoins        string    `json:"no_of_coins"`
}

func (server *Server) createTransaction(ctx *gin.Context) {
	var req createTransactionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateTransactionParams{
		AccountID:        req.AccountID,
		CoinID:           req.CoinID,
		CoinName:         req.CoinName,
		Symbol:           req.Symbol,
		Type:             req.Type,
		Amount:           req.Amount,
		TimeTransacted:   req.TimeTransacted,
		PricePurchasedAt: req.PricePurchasedAt,
		NoOfCoins:        req.NoOfCoins,
	}
	transaction, err := server.store.CreateTransaction(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, transaction)
}