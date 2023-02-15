package api

import (
	"net/http"
	db "server/db/sqlc"

	"github.com/gin-gonic/gin"
)

type createPortfolioRequest struct {
	Name string `json:"name"`
	AccountID int64 `json:"account_id"`
}

func (server *Server) createPortfolio(ctx *gin.Context) {
	var req createPortfolioRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreatePortfolioParams{
		Name: req.Name,
		AccountID: req.AccountID,
	}
	portfolio, err := server.store.CreatePortfolio(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, portfolio)
}