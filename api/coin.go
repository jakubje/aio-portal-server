package api

import (
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/jakub/aioportal/server/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createCoinRequest struct {
	CoinID            string   `json:"coin_id"`
	Name              string   `json:"name"`
	Price             string   `json:"price"`
	MarketCap         string   `json:"market_cap"`
	CirculatingSupply string   `json:"circulating_supply"`
	TotalSupply       string   `json:"total_supply"`
	MaxSupply         string   `json:"max_supply"`
	Rank              string   `json:"rank"`
	Volume            string   `json:"volume"`
	ImageUrl          string   `json:"image_url"`
	Description       string   `json:"description"`
	Website           string   `json:"website"`
	SocialMediaLinks  []string `json:"social_media_links"`
}

func (server *Server) createCoin(ctx *gin.Context) {
	var req createCoinRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateCoinParams{
		CoinID:            req.CoinID,
		Name:              req.Name,
		Price:             req.Price,
		MarketCap:         req.MarketCap,
		CirculatingSupply: req.CirculatingSupply,
		TotalSupply:       req.TotalSupply,
		MaxSupply:         req.MaxSupply,
		Rank:              req.Rank,
		Volume:            req.Volume,
		ImageUrl:          req.ImageUrl,
		Description:       req.Description,
		Website:           req.Website,
		SocialMediaLinks:  req.SocialMediaLinks,
	}

	coin, err := server.store.CreateCoin(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, coin)
}

type updateCoinRequest struct {
	CoinID            string   `json:"coin_id"`
	Name              string   `json:"name"`
	Price             string   `json:"price"`
	MarketCap         string   `json:"market_cap"`
	CirculatingSupply string   `json:"circulating_supply"`
	TotalSupply       string   `json:"total_supply"`
	MaxSupply         string   `json:"max_supply"`
	Rank              string   `json:"rank"`
	Volume            string   `json:"volume"`
	ImageUrl          string   `json:"image_url"`
	Description       string   `json:"description"`
	Website           string   `json:"website"`
	SocialMediaLinks  []string `json:"social_media_links"`
}

func (server *Server) updateCoin(ctx *gin.Context) {
	var req updateCoinRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateCoinParams{
		Name: pgtype.Text{
			String: req.Name,
			Valid:  true,
		},
		Price: pgtype.Text{
			String: req.Price,
			Valid:  true,
		},
		MarketCap: pgtype.Text{
			String: req.MarketCap,
			Valid:  true,
		},
		CirculatingSupply: pgtype.Text{
			String: req.CirculatingSupply,
			Valid:  true,
		},
		TotalSupply: pgtype.Text{
			String: req.TotalSupply,
			Valid:  true,
		},
		MaxSupply: pgtype.Text{
			String: req.MaxSupply,
			Valid:  true,
		},
		Rank: pgtype.Text{
			String: req.Rank,
			Valid:  true,
		},
		Volume: pgtype.Text{
			String: req.Volume,
			Valid:  true,
		},
		ImageUrl: pgtype.Text{
			String: req.ImageUrl,
			Valid:  true,
		},
		Description: pgtype.Text{
			String: req.Description,
			Valid:  true,
		},
		Website: pgtype.Text{
			String: req.Website,
			Valid:  true,
		},
		SocialMediaLinks: nil,
		CoinID:           req.CoinID,
	}
	coin, err := server.store.UpdateCoin(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, coin)
}
