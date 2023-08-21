package api

import (
	"errors"
	db "github.com/jakub/aioportal/server/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddWatchlistCoinRequest struct {
	WatchlistID int64  `json:"watchlist_id"`
	CoinID      string `json:"coin_id"`
}

func (server *Server) addWatchlistCoin(ctx *gin.Context) {
	var req AddWatchlistCoinRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.AddWatchlistCoinParams{
		WatchlistID: req.WatchlistID,
		CoinID:      req.CoinID,
	}
	_, err := server.store.AddWatchlistCoin(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": " coin added"})
}

type removeWatchlistCoin struct {
	CoinID      string `uri:"coin_id" binding:"required,min=1"`
	WatchlistID int64  `uri:"watchlist_id" binding:"required,min=1"`
}

func (server *Server) removeWatchlistCoin(ctx *gin.Context) {
	var req removeWatchlistCoin
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.RemoveWatchlistCoinParams{
		WatchlistID: req.WatchlistID,
		CoinID:      req.CoinID,
	}
	err := server.store.RemoveWatchlistCoin(ctx, arg)

	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "watchlist coin deleted"})
}
