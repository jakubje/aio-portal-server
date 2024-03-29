package api

import (
	"errors"
	"github.com/jakub/aioportal/server/token"
	"net/http"

	db "github.com/jakub/aioportal/server/db/sqlc"

	"github.com/gin-gonic/gin"
)

type watchListRequestResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type watchListsResponse struct {
	Total      int64                      `json:"total"`
	Watchlists []watchListRequestResponse `json:"watchlists"`
}

func (server *Server) createWatchlist(ctx *gin.Context) {
	var req watchListRequestResponse
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.CreateWatchlistParams{
		Name:      req.Name,
		AccountID: authPayload.AccountId,
	}
	watchlist, err := server.store.CreateWatchlist(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := watchListRequestResponse{
		ID:   watchlist.ID,
		Name: watchlist.Name,
	}
	ctx.JSON(http.StatusOK, rsp)
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
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.DeleteWatchlistParams{
		ID:        req.ID,
		AccountID: authPayload.AccountId,
	}

	err := server.store.DeleteWatchlist(ctx, arg)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
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
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.GetWatchlistParams{
		ID:        req.ID,
		AccountID: authPayload.AccountId,
	}

	watchlist, err := server.store.GetWatchlist(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := watchListRequestResponse{
		ID:   watchlist.ID,
		Name: watchlist.Name,
	}
	ctx.JSON(http.StatusOK, rsp)
}

func (server *Server) listWatchlists(ctx *gin.Context) {
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	watchlists, err := server.store.ListWatchlists(ctx, authPayload.AccountId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp := watchListsResponse{}
	for _, watchList := range watchlists {
		rsp := watchListRequestResponse{
			ID:   watchList.ID,
			Name: watchList.Name,
		}
		resp.Watchlists = append(resp.Watchlists, rsp)
	}
	resp.Total = int64(len(watchlists))

	ctx.JSON(http.StatusOK, resp)
}

type updateWatchlistRequest struct {
	ID   int64  `json:"id" binding:"required"`
	Name string `json:"name"`
}

func (server *Server) updateWatchlist(ctx *gin.Context) {
	var req updateWatchlistRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.UpdateWatchlistParams{
		AccountID: authPayload.AccountId,
		ID:        req.ID,
		Name:      req.Name,
	}

	watchlist, err := server.store.UpdateWatchlist(ctx, arg)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := watchListRequestResponse{
		ID:   watchlist.ID,
		Name: watchlist.Name,
	}

	ctx.JSON(http.StatusOK, rsp)
}
