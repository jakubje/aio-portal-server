// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: coin.sql

package db

import (
	"context"
)

const createCoin = `-- name: CreateCoin :one
INSERT INTO coins (
    portfolio_id, coin_name, coin_symbol, amount, no_of_coins
) VALUES (
             $1, $2, $3, $4, $5
         )
RETURNING id, portfolio_id, coin_name, coin_symbol, amount, time_created, no_of_coins
`

type CreateCoinParams struct {
	PortfolioID int64  `json:"portfolio_id"`
	CoinName    string `json:"coin_name"`
	CoinSymbol  string `json:"coin_symbol"`
	Amount      int32  `json:"amount"`
	NoOfCoins   string `json:"no_of_coins"`
}

func (q *Queries) CreateCoin(ctx context.Context, arg CreateCoinParams) (Coin, error) {
	row := q.db.QueryRowContext(ctx, createCoin,
		arg.PortfolioID,
		arg.CoinName,
		arg.CoinSymbol,
		arg.Amount,
		arg.NoOfCoins,
	)
	var i Coin
	err := row.Scan(
		&i.ID,
		&i.PortfolioID,
		&i.CoinName,
		&i.CoinSymbol,
		&i.Amount,
		&i.TimeCreated,
		&i.NoOfCoins,
	)
	return i, err
}

const deleteCoin = `-- name: DeleteCoin :exec
DELETE FROM coins
WHERE id = $1
`

func (q *Queries) DeleteCoin(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteCoin, id)
	return err
}

const getCoin = `-- name: GetCoin :one
SELECT id, portfolio_id, coin_name, coin_symbol, amount, time_created, no_of_coins FROM coins
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetCoin(ctx context.Context, id int64) (Coin, error) {
	row := q.db.QueryRowContext(ctx, getCoin, id)
	var i Coin
	err := row.Scan(
		&i.ID,
		&i.PortfolioID,
		&i.CoinName,
		&i.CoinSymbol,
		&i.Amount,
		&i.TimeCreated,
		&i.NoOfCoins,
	)
	return i, err
}

const listCoins = `-- name: ListCoins :many
SELECT id, portfolio_id, coin_name, coin_symbol, amount, time_created, no_of_coins FROM coins
WHERE portfolio_id = $1
`

func (q *Queries) ListCoins(ctx context.Context, portfolioID int64) ([]Coin, error) {
	rows, err := q.db.QueryContext(ctx, listCoins, portfolioID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Coin{}
	for rows.Next() {
		var i Coin
		if err := rows.Scan(
			&i.ID,
			&i.PortfolioID,
			&i.CoinName,
			&i.CoinSymbol,
			&i.Amount,
			&i.TimeCreated,
			&i.NoOfCoins,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCoin = `-- name: UpdateCoin :one
UPDATE coins
set amount = $2,
    no_of_coins = $3
WHERE id = $1
RETURNING id, portfolio_id, coin_name, coin_symbol, amount, time_created, no_of_coins
`

type UpdateCoinParams struct {
	ID        int64  `json:"id"`
	Amount    int32  `json:"amount"`
	NoOfCoins string `json:"no_of_coins"`
}

func (q *Queries) UpdateCoin(ctx context.Context, arg UpdateCoinParams) (Coin, error) {
	row := q.db.QueryRowContext(ctx, updateCoin, arg.ID, arg.Amount, arg.NoOfCoins)
	var i Coin
	err := row.Scan(
		&i.ID,
		&i.PortfolioID,
		&i.CoinName,
		&i.CoinSymbol,
		&i.Amount,
		&i.TimeCreated,
		&i.NoOfCoins,
	)
	return i, err
}
