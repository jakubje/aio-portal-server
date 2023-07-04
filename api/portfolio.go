package api

import (
	"database/sql"
	db "github.com/jakub/aioportal/server/db/sqlc"
	"github.com/jakub/aioportal/server/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createPortfolioRequest struct {
	Name string `json:"name"`
}

func (server *Server) createPortfolio(ctx *gin.Context) {
	var req createPortfolioRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.CreatePortfolioParams{
		Name:      req.Name,
		AccountID: authPayload.AccountId,
	}

	portfolio, err := server.store.CreatePortfolio(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, portfolio)
}

type getPortfolioRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type portfolioResponse struct {
	Name       string `json:"name"`
	Holdings   int32  `json:"holdings"`
	Change24H  int32  `json:"change_24h"`
	ProfitLoss int32  `json:"profit_loss"`
}

func (server *Server) getPortfolio(ctx *gin.Context) {
	var req getPortfolioRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.GetPortfolioParams{
		ID:        req.ID,
		AccountID: authPayload.AccountId,
	}

	portfolio, err := server.store.GetPortfolio(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	rsp := portfolioResponse{
		Name:       portfolio.Name,
		Holdings:   portfolio.Holdings,
		Change24H:  portfolio.Change24h,
		ProfitLoss: portfolio.ProfitLoss,
	}
	ctx.JSON(http.StatusOK, rsp)
}

type listPortfoliosRequest struct {
	AccountID int64 `uri:"account_id" binding:"required,min=1"`
}

func (server *Server) listPortfolios(ctx *gin.Context) {
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	portfolios, err := server.store.ListPortforlios(ctx, authPayload.AccountId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, portfolios)
}

type updatePortfolioRequest struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (server *Server) updatePortfolio(ctx *gin.Context) {
	var req updatePortfolioRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.UpdatePortfolioParams{
		ID:        req.ID,
		Name:      req.Name,
		AccountID: authPayload.AccountId,
	}

	portfolio, err := server.store.UpdatePortfolio(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, portfolio)
}

type deletePortfolioRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deletePortfolio(ctx *gin.Context) {
	var req deletePortfolioRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.DeletePortfolioParams{
		ID:        req.ID,
		AccountID: authPayload.AccountId,
	}

	err := server.store.DeletePortfolio(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "portfolio deleted"})
}

type getRollUpByCoinByPortfolioRequest struct {
	PortfolioID int64 `uri:"portfolio_id" binding:"required,min=1"`
}

func (server *Server) getRollUpByCoinByPortfolio(ctx *gin.Context) {
	var req getRollUpByCoinByPortfolioRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.GetRollUpByCoinByPortfolioParams{
		PortfolioID: req.PortfolioID,
		AccountID:   authPayload.AccountId,
	}

	rollup, err := server.store.GetRollUpByCoinByPortfolio(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, rollup)
}
