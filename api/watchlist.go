package api

import (
	"database/sql"
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

type deleteWatchlist struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteWatchlist(ctx *gin.Context) {
	var req deleteWatchlist
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	err := server.store.DeleteWatchlist(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "watchlist deleted"})
}


type getWatchlistRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getWatchlist(ctx *gin.Context) {
	var req getWatchlistRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	watchlist, err := server.store.GetWatchlist(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, watchlist)
}

type listWatchlistsRequest struct {
	AccountID int64 `uri:"account_id" binding:"required,min=1"`
}

func (server *Server) listWatchlists(ctx *gin.Context) {
	var req listWatchlistsRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	watchlists, err := server.store.ListWatchlists(ctx, req.AccountID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, watchlists)
}


type updateWatchlistRequest struct {
	ID        int64  `uri:"id" binding:"required,min=1"`
	Name      string `json:"name"`
}

func (server *Server) updateWatchlist(ctx *gin.Context) {
	var req updateWatchlistRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateWatchlistParams{
		ID:        req.ID,
		Name:      req.Name,
	}
	watchlist, err := server.store.UpdateWatchlist(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, watchlist)
}

