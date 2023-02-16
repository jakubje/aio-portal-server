package api

import (
	"database/sql"
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

type getTransactionRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getTransaction(ctx *gin.Context) {
	var req getTransactionRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	transaction, err := server.store.GetTransaction(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, transaction)
}

type listTransactionsByAccountRequest struct {
	ID     int64 `form:"id" binding:"required,min=1"`
	Limit  int32 `form:"limit,default=100" binding:"max=100"`
	Offset int32 `form:"offset,default=0"`
}

func (server *Server) listTransactions(ctx *gin.Context) {
	var req listTransactionsByAccountRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListTransactionsByAccountParams{
		AccountID:     req.ID,
		Limit:  req.Limit,
		Offset: req.Offset,
	}

	transactions, err := server.store.ListTransactionsByAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, transactions)
}

type listTransactionsByAccountByCoinRequest struct {
	AccountID int64 `json:"account_id" binding:"required,min=1"`
	CoinID    int64 `json:"coin" binding:"required"`
	Limit     int32  `json:"limit,default=10" binding:"max=100"`
	Offset    int32  `json:"offset,default=0"`
}


func (server *Server) listTransactionsByAccountByCoin(ctx *gin.Context) {
	var request listTransactionsByAccountByCoinRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListTransactionsByAccountByCoinParams{
		AccountID: request.AccountID,
		CoinID:      request.CoinID,
		Limit:     request.Limit,
		Offset:    request.Offset,
	}

	transactions, err := server.store.ListTransactionsByAccountByCoin(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, transactions)
}

type deleteTransactionReqiest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteTransaction(ctx *gin.Context) {
	var req deleteTransactionReqiest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteTransaction(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "transaction deleted"})
}

