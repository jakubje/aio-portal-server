package api

import (
	"database/sql"
	db "github.com/jakub/aioportal/server/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createPortfolioRequest struct {
	Name      string `json:"name"`
	AccountID int64  `json:"account_id"`
}

func (server *Server) createPortfolio(ctx *gin.Context) {
	var req createPortfolioRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreatePortfolioParams{
		Name:      req.Name,
		AccountID: req.AccountID,
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

func (server *Server) getPortfolio(ctx *gin.Context) {
	var req getPortfolioRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	portfolio, err := server.store.GetPortfolio(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, portfolio)
}

type listPortfoliosRequest struct {
	AccountID int64 `uri:"account_id" binding:"required,min=1"`
}

func (server *Server) listPortfolios(ctx *gin.Context) {
	var req listPortfoliosRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	portfolios, err := server.store.ListPortforlios(ctx, req.AccountID)
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

	err := server.store.DeletePortfolio(ctx, req.ID)
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
