package api

import (
	"database/sql"
	db "github.com/jakub/aioportal/server/db/sqlc"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type createTransactionRequest struct {
	AccountID    int64   `json:"account_id"`
	PortfolioID  int64   `json:"portfolio_id"`
	CoinName     string  `json:"coin_name"`
	Symbol       string  `json:"symbol"`
	Type         int32   `json:"type"`
	Quantity     float64 `json:"quantity"`
	PricePerCoin float64 `json:"price_per_coin"`
}

func (server *Server) createTransaction(ctx *gin.Context) {
	var req createTransactionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateTransactionParams{
		AccountID:      req.AccountID,
		PortfolioID:    req.PortfolioID,
		Symbol:         req.Symbol,
		Type:           req.Type,
		PricePerCoin:   req.PricePerCoin,
		Quantity:       req.Quantity,
		TimeTransacted: time.Now(),
		TimeCreated:    time.Now(),
	}
	transaction, err := server.store.CreateTransaction(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, transaction)
}

type getTransactionRequest struct {
	ID string `uri:"id" binding:"required,min=1"`
}

func (server *Server) getTransaction(ctx *gin.Context) {
	var req getTransactionRequest
	log.Printf("getTransactionRequest: %v", req)
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	parsedId, _ := uuid.Parse(req.ID)
	transaction, err := server.store.GetTransaction(ctx, parsedId)
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
		AccountID: req.ID,
		Limit:     req.Limit,
		Offset:    req.Offset,
	}

	transactions, err := server.store.ListTransactionsByAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, transactions)
}

type listTransactionsByAccountByCoinRequest struct {
	AccountID int64  `json:"account_id" binding:"required,min=1"`
	Symbol    string `json:"symbol" binding:"required"`
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
		Symbol:    request.Symbol,
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
	ID string `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteTransaction(ctx *gin.Context) {
	var req deleteTransactionReqiest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	parsedId, _ := uuid.Parse(req.ID)
	err := server.store.DeleteTransaction(ctx, parsedId)
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
