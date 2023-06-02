package api

import (
	db "github.com/jakub/aioportal/server/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type addFootballRequest struct {
	AccountID int64  `json:"account_id"`
	Team      string `json:"team"`
	League    string `json:"league"`
	Country   string `json:"country"`
}

func (server *Server) addFootball(ctx *gin.Context) {
	var req addFootballRequest
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateFootballParams{
		AccountID: req.AccountID,
		Team:      req.Team,
		League:    req.League,
		Country:   req.Country,
	}
	football, err := server.store.CreateFootball(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, football)
}

func MarshalJSON(req addFootballRequest) {
	panic("unimplemented")
}

type getFootballRequest struct {
	AccountID int64 `uri:"id" binding:"required"`
}

func (server *Server) getFootball(ctx *gin.Context) {
	var req getFootballRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	football, err := server.store.GetFootball(ctx, req.AccountID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, football)
}

type updateFootballParams struct {
	AccountID int64  `json:"account_id"`
	Team      string `json:"team"`
	League    string `json:"league"`
	Country   string `json:"country"`
}

func (server *Server) updateFootball(ctx *gin.Context) {
	var req updateFootballParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateFootballParams{
		AccountID: req.AccountID,
		Team:      req.Team,
		League:    req.League,
		Country:   req.Country,
	}
	football, err := server.store.UpdateFootball(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, football)
}
