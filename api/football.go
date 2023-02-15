package api

import (
	"database/sql"
	"net/http"
	db "server/db/sqlc"

	"github.com/gin-gonic/gin"
)

type addFootballRequest struct {
	AccountID int64          `json:"account_id"`
	Team      sql.NullString `json:"team"`
	League    sql.NullString `json:"league"`
	Country   sql.NullString `json:"country"`
}

func (server *Server) addFootball(ctx *gin.Context) {
	var req addFootballRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
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