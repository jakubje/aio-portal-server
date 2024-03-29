package api

import (
	db "github.com/jakub/aioportal/server/db/sqlc"
	"github.com/jakub/aioportal/server/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

type footballRequestResponse struct {
	Team    string `json:"team"`
	League  string `json:"league"`
	Country string `json:"country"`
}

func (server *Server) addFootball(ctx *gin.Context) {
	var req footballRequestResponse
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.CreateFootballParams{
		AccountID: authPayload.AccountId,
		Team:      req.Team,
		League:    req.League,
		Country:   req.Country,
	}

	football, err := server.store.CreateFootball(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	rsp := footballRequestResponse{
		Team:    football.Team,
		League:  football.League,
		Country: football.Country,
	}
	ctx.JSON(http.StatusOK, rsp)
}

func (server *Server) getFootball(ctx *gin.Context) {

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	football, err := server.store.GetFootball(ctx, authPayload.AccountId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := footballRequestResponse{
		Team:    football.Team,
		League:  football.League,
		Country: football.Country,
	}

	ctx.JSON(http.StatusOK, rsp)
}

func (server *Server) updateFootball(ctx *gin.Context) {
	var req footballRequestResponse
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.UpdateFootballParams{
		AccountID: authPayload.AccountId,
		Team:      req.Team,
		League:    req.League,
		Country:   req.Country,
	}
	football, err := server.store.UpdateFootball(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := footballRequestResponse{
		Team:    football.Team,
		League:  football.League,
		Country: football.Country,
	}

	ctx.JSON(http.StatusOK, rsp)
}
