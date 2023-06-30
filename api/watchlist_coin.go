package api

import (
	"database/sql"
	db "github.com/jakub/aioportal/server/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createWatchlistCoinRequest struct {
	WatchlistID int64  `json:"watchlist_id"`
	Name        string `json:"name"`
	Symbol      string `json:"symbol"`
	Rank        int16  `json:"rank"`
}

type watchListCoinResponse struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	Rank   int16  `json:"rank"`
}

type watchListCoinsResponse struct {
	Total int64                   `json:"total"`
	Coins []watchListCoinResponse `json:"coins"`
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

	rsp := watchListCoinResponse{
		Name:   coin.Name,
		Symbol: coin.Symbol,
		Rank:   coin.Rank,
	}
	ctx.JSON(http.StatusOK, rsp)
}

type deleteWachlistCoin struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteWachlistCoin(ctx *gin.Context) {
	var req deleteWachlistCoin
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	err := server.store.DeleteWatchlistCoin(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "watchlist coin deleted"})
}

type getWatchlistCoinRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getWatchlistCoin(ctx *gin.Context) {
	var req getWatchlistCoinRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	coin, err := server.store.GetWatchlistCoin(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := watchListCoinResponse{
		Name:   coin.Name,
		Symbol: coin.Symbol,
		Rank:   coin.Rank,
	}
	ctx.JSON(http.StatusOK, rsp)
}

type listWatchlistCoins struct {
	WatchlistID int64 `uri:"watchlist_id" binding:"required,min=1"`
}

func (server *Server) listWatchlistCoins(ctx *gin.Context) {
	var req listWatchlistCoins
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	coins, err := server.store.ListWatchlistsCoins(ctx, req.WatchlistID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp := watchListCoinsResponse{}
	for _, coin := range coins {
		rsp := watchListCoinResponse{
			ID:     coin.ID,
			Name:   coin.Name,
			Symbol: coin.Symbol,
			Rank:   coin.Rank,
		}
		resp.Coins = append(resp.Coins, rsp)
	}
	resp.Total = int64(len(coins))

	ctx.JSON(http.StatusOK, resp)
}
