package api

import (
	"database/sql"
	"errors"
	db "github.com/jakub/aioportal/server/db/sqlc"
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

	accountId, exists := server.ctx.Get("accountId")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("missing account id")))
		return
	}

	arg := db.CreatePortfolioParams{
		Name:      req.Name,
		AccountID: accountId.(int64),
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

	accountId, exists := server.ctx.Get("accountId")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("missing account id")))
		return
	}
	arg := db.GetPortfolioParams{
		ID:        req.ID,
		AccountID: accountId.(int64),
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
	//var req listPortfoliosRequest
	//if err := ctx.ShouldBindUri(&req); err != nil {
	//	ctx.JSON(http.StatusBadRequest, errorResponse(err))
	//	return
	//}
	accountId, exists := server.ctx.Get("accountId")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("missing account id")))
		return
	}
	portfolios, err := server.store.ListPortforlios(ctx, accountId.(int64))
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

	arg := db.UpdatePortfolioParams{
		ID:   req.ID,
		Name: req.Name,
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

	accountId, exists := server.ctx.Get("accountId")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("missing account id")))
		return
	}
	arg := db.DeletePortfolioParams{
		ID:        req.ID,
		AccountID: accountId.(int64),
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

	rollup, err := server.store.GetRollUpByCoinByPortfolio(ctx, req.PortfolioID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, rollup)
}
