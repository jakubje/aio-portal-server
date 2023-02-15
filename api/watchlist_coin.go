package api

import (
	"net/http"
	db "server/db/sqlc"

	"github.com/gin-gonic/gin"
)

type createWatchlistCoinRequest struct {
	WatchlistID int64  `json:"watchlist_id"`
	Name        string `json:"name"`
	Symbol      string `json:"symbol"`
	Rank        int16  `json:"rank"`
}

func (server *Server) createWatchlistCoin(ctx *gin.Context) {
	var req createWatchlistCoinRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateWatchlistCoinsParams{
		WatchlistID: req.WatchlistID,
		Name:        req.Name,
		Symbol:      req.Symbol,
		Rank:        req.Rank,
	}
	coin, err := server.store.CreateWatchlistCoins(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, coin)
}
