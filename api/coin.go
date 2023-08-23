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
	Price             float64  `json:"price"`
	MarketCap         int64    `json:"market_cap"`
	CirculatingSupply int64    `json:"circulating_supply"`
	TotalSupply       int64    `json:"total_supply"`
	MaxSupply         int64    `json:"max_supply"`
	Rank              int32    `json:"rank"`
	Volume            int64    `json:"volume"`
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
	Price             float64  `json:"price"`
	MarketCap         int64    `json:"market_cap"`
	CirculatingSupply int64    `json:"circulating_supply"`
	TotalSupply       int64    `json:"total_supply"`
	MaxSupply         int64    `json:"max_supply"`
	Rank              int32    `json:"rank"`
	Volume            int64    `json:"volume"`
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
		Price: pgtype.Float8{
			Float64: req.Price,
			Valid:   true,
		},
		MarketCap: pgtype.Int8{
			Int64: req.MarketCap,
			Valid: true,
		},
		CirculatingSupply: pgtype.Int8{
			Int64: req.CirculatingSupply,
			Valid: true,
		},
		TotalSupply: pgtype.Int8{
			Int64: req.TotalSupply,
			Valid: true,
		},
		MaxSupply: pgtype.Int8{
			Int64: req.MaxSupply,
			Valid: true,
		},
		Rank: pgtype.Int4{
			Int32: req.Rank,
			Valid: true,
		},
		Volume: pgtype.Int8{
			Int64: req.Volume,
			Valid: true,
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
