package api

import (
	"net/http"
	db "server/db/sqlc"

	"github.com/gin-gonic/gin"
)


type createWatchlistRequest struct {
	Name      string `json:"name"`
	AccountID int64  `json:"account_id"`
}

func (server *Server) createWatchlist(ctx *gin.Context) {
	var req createWatchlistRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateWatchlistParams{
		Name:      req.Name,
		AccountID: req.AccountID,
	}
	watchlist, err := server.store.CreateWatchlist(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, watchlist)
}