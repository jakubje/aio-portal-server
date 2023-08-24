package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/google/uuid"
	db "github.com/jakub/aioportal/server/db/sqlc"
	"github.com/jakub/aioportal/server/token"

	"github.com/gin-gonic/gin"
)

type createTransactionRequest struct {
	AccountID    int64   `json:"account_id"`
	PortfolioID  int64   `json:"portfolio_id"`
	Symbol       string  `json:"symbol"`
	Type         int32   `json:"type"`
	Quantity     float64 `json:"quantity"`
	PricePerCoin float64 `json:"price_per_coin"`
}

type transactionResponse struct {
	ID           uuid.UUID `json:"id"`
	PortfolioID  int64     `json:"portfolio_id"`
	Symbol       string    `json:"symbol"`
	Type         int32     `json:"type"`
	Quantity     float64   `json:"quantity"`
	PricePerCoin float64   `json:"price_per_coin"`
	TimeCreated  time.Time `json:"time_created"`
}

type transactionsResponse struct {
	Total        int64                 `json:"total"`
	Transactions []transactionResponse `json:"transactions"`
}

func (server *Server) createTransaction(ctx *gin.Context) {
	var req createTransactionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.CreateTransactionParams{
		AccountID:      authPayload.AccountId,
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

	rsp := transactionResponse{
		ID:           transaction.ID,
		PortfolioID:  transaction.PortfolioID,
		Symbol:       transaction.Symbol,
		Type:         transaction.Type,
		Quantity:     transaction.Quantity,
		PricePerCoin: transaction.PricePerCoin,
	}
	ctx.JSON(http.StatusOK, rsp)
}

type getTransactionRequest struct {
	ID string `uri:"id" binding:"required,min=1"`
}

func (server *Server) getTransaction(ctx *gin.Context) {
	var req getTransactionRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	parsedId, _ := uuid.Parse(req.ID)
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := db.GetTransactionParams{
		ID:        parsedId,
		AccountID: authPayload.AccountId,
	}

	transaction, err := server.store.GetTransaction(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := transactionResponse{
		ID:           transaction.ID,
		PortfolioID:  transaction.PortfolioID,
		Symbol:       transaction.Symbol,
		Type:         transaction.Type,
		Quantity:     transaction.Quantity,
		PricePerCoin: transaction.PricePerCoin,
	}

	ctx.JSON(http.StatusOK, rsp)
}

type listTransactionsByAccountRequest struct {
	Limit  int32 `form:"limit,default=100" binding:"max=100"`
	Offset int32 `form:"offset,default=0"`
}

func (server *Server) listTransactions(ctx *gin.Context) {
	var req listTransactionsByAccountRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.ListTransactionsByAccountParams{
		AccountID: authPayload.AccountId,
		Limit:     req.Limit,
		Offset:    req.Offset,
	}

	transactions, err := server.store.ListTransactionsByAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp := transactionsResponse{}
	for _, transaction := range transactions {
		rsp := transactionResponse{
			ID:           transaction.ID,
			PortfolioID:  transaction.PortfolioID,
			Symbol:       transaction.Symbol,
			Type:         transaction.Type,
			Quantity:     transaction.Quantity,
			PricePerCoin: transaction.PricePerCoin,
		}
		resp.Transactions = append(resp.Transactions, rsp)
	}
	resp.Total = int64(len(transactions))

	ctx.JSON(http.StatusOK, resp)
}

type listTransactionsByAccountByCoinRequest struct {
	Symbol string `json:"symbol" binding:"required"`
	Limit  int32  `json:"limit,default=10" binding:"max=100"`
	Offset int32  `json:"offset,default=0"`
}

func (server *Server) listTransactionsByAccountByCoin(ctx *gin.Context) {
	var request listTransactionsByAccountByCoinRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.ListTransactionsByAccountByCoinParams{
		AccountID: authPayload.AccountId,
		Symbol:    request.Symbol,
		Limit:     request.Limit,
		Offset:    request.Offset,
	}

	transactions, err := server.store.ListTransactionsByAccountByCoin(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp := transactionsResponse{}
	for _, transaction := range transactions {
		rsp := transactionResponse{
			ID:           transaction.ID,
			PortfolioID:  transaction.PortfolioID,
			Symbol:       transaction.Symbol,
			Type:         transaction.Type,
			Quantity:     transaction.Quantity,
			PricePerCoin: transaction.PricePerCoin,
		}
		resp.Transactions = append(resp.Transactions, rsp)
	}
	resp.Total = int64(len(transactions))

	ctx.JSON(http.StatusOK, resp)
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
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.DeleteTransactionParams{
		ID:        parsedId,
		AccountID: authPayload.AccountId,
	}

	err := server.store.DeleteTransaction(ctx, arg)
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
