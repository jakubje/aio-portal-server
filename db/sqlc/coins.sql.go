// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: coins.sql

package db

import (
	"context"
	"time"
)

const createCoin = `-- name: CreateCoin :one
INSERT INTO coins (
    coin_id, coin_uuid, name, price, market_cap, number_of_markets, number_of_exchanges, approved_supply, circulating_supply, total_supply, max_supply, rank, volume, daily_change, image_url, description, all_time_high, tags, website, social_media_links, updated_at
) VALUES (
             $1, $2, $3, $4, $5, $6, $7, $8 ,$9 ,$10 ,$11 ,$12 ,$13 ,$14, $15, $16, $17, $18, $19, $20, $21
         ) ON CONFLICT (coin_id) DO UPDATE
SET
    coin_uuid = $2,
    name = $3,
    price = $4,
    market_cap = $5,
    number_of_markets = $6,
    number_of_exchanges = $7,
    approved_supply = $8,
    circulating_supply = $9,
    total_supply =  $10,
    max_supply = $11,
    rank = $12,
    volume = $13,
    daily_change = $14,
    image_url = $15,
    description = $16,
    all_time_high = $17,
    tags = $18,
    website = $19,
    social_media_links = $20,
    updated_at = $21
RETURNING coin_id, coin_uuid, name, price, market_cap, number_of_markets, number_of_exchanges, approved_supply, circulating_supply, total_supply, max_supply, rank, volume, daily_change, image_url, description, all_time_high, tags, website, social_media_links, created_at, updated_at
`

type CreateCoinParams struct {
	CoinID            string    `json:"coin_id"`
	CoinUuid          string    `json:"coin_uuid"`
	Name              string    `json:"name"`
	Price             string    `json:"price"`
	MarketCap         string    `json:"market_cap"`
	NumberOfMarkets   int32     `json:"number_of_markets"`
	NumberOfExchanges int32     `json:"number_of_exchanges"`
	ApprovedSupply    bool      `json:"approved_supply"`
	CirculatingSupply string    `json:"circulating_supply"`
	TotalSupply       string    `json:"total_supply"`
	MaxSupply         string    `json:"max_supply"`
	Rank              int32     `json:"rank"`
	Volume            string    `json:"volume"`
	DailyChange       string    `json:"daily_change"`
	ImageUrl          string    `json:"image_url"`
	Description       string    `json:"description"`
	AllTimeHigh       string    `json:"all_time_high"`
	Tags              []string  `json:"tags"`
	Website           string    `json:"website"`
	SocialMediaLinks  []string  `json:"social_media_links"`
	UpdatedAt         time.Time `json:"updated_at"`
}

func (q *Queries) CreateCoin(ctx context.Context, arg CreateCoinParams) (Coin, error) {
	row := q.db.QueryRow(ctx, createCoin,
		arg.CoinID,
		arg.CoinUuid,
		arg.Name,
		arg.Price,
		arg.MarketCap,
		arg.NumberOfMarkets,
		arg.NumberOfExchanges,
		arg.ApprovedSupply,
		arg.CirculatingSupply,
		arg.TotalSupply,
		arg.MaxSupply,
		arg.Rank,
		arg.Volume,
		arg.DailyChange,
		arg.ImageUrl,
		arg.Description,
		arg.AllTimeHigh,
		arg.Tags,
		arg.Website,
		arg.SocialMediaLinks,
		arg.UpdatedAt,
	)
	var i Coin
	err := row.Scan(
		&i.CoinID,
		&i.CoinUuid,
		&i.Name,
		&i.Price,
		&i.MarketCap,
		&i.NumberOfMarkets,
		&i.NumberOfExchanges,
		&i.ApprovedSupply,
		&i.CirculatingSupply,
		&i.TotalSupply,
		&i.MaxSupply,
		&i.Rank,
		&i.Volume,
		&i.DailyChange,
		&i.ImageUrl,
		&i.Description,
		&i.AllTimeHigh,
		&i.Tags,
		&i.Website,
		&i.SocialMediaLinks,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getCoin = `-- name: GetCoin :one
SELECT coin_id, coin_uuid, name, price, market_cap, number_of_markets, number_of_exchanges, approved_supply, circulating_supply, total_supply, max_supply, rank, volume, daily_change, image_url, description, all_time_high, tags, website, social_media_links, created_at, updated_at FROM coins
WHERE coin_id = $1
LIMIT 1
`

func (q *Queries) GetCoin(ctx context.Context, coinID string) (Coin, error) {
	row := q.db.QueryRow(ctx, getCoin, coinID)
	var i Coin
	err := row.Scan(
		&i.CoinID,
		&i.CoinUuid,
		&i.Name,
		&i.Price,
		&i.MarketCap,
		&i.NumberOfMarkets,
		&i.NumberOfExchanges,
		&i.ApprovedSupply,
		&i.CirculatingSupply,
		&i.TotalSupply,
		&i.MaxSupply,
		&i.Rank,
		&i.Volume,
		&i.DailyChange,
		&i.ImageUrl,
		&i.Description,
		&i.AllTimeHigh,
		&i.Tags,
		&i.Website,
		&i.SocialMediaLinks,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listCoins = `-- name: ListCoins :many
SELECT coin_id, coin_uuid, name, price, market_cap, number_of_markets, number_of_exchanges, approved_supply, circulating_supply, total_supply, max_supply, rank, volume, daily_change, image_url, description, all_time_high, tags, website, social_media_links, created_at, updated_at FROM coins
ORDER BY rank
LIMIT $1
OFFSET $2
`

type ListCoinsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListCoins(ctx context.Context, arg ListCoinsParams) ([]Coin, error) {
	rows, err := q.db.Query(ctx, listCoins, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Coin{}
	for rows.Next() {
		var i Coin
		if err := rows.Scan(
			&i.CoinID,
			&i.CoinUuid,
			&i.Name,
			&i.Price,
			&i.MarketCap,
			&i.NumberOfMarkets,
			&i.NumberOfExchanges,
			&i.ApprovedSupply,
			&i.CirculatingSupply,
			&i.TotalSupply,
			&i.MaxSupply,
			&i.Rank,
			&i.Volume,
			&i.DailyChange,
			&i.ImageUrl,
			&i.Description,
			&i.AllTimeHigh,
			&i.Tags,
			&i.Website,
			&i.SocialMediaLinks,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}