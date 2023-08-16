// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: watchlist_coin.sql

package db

import (
	"context"
)

const createWatchlistCoins = `-- name: CreateWatchlistCoins :one
INSERT INTO watchlist_coins (
    watchlist_id, name, symbol, rank
) VALUES (
             $1, $2, $3, $4
         )
RETURNING id, watchlist_id, name, symbol, rank
`

type CreateWatchlistCoinsParams struct {
	WatchlistID int64  `json:"watchlist_id"`
	Name        string `json:"name"`
	Symbol      string `json:"symbol"`
	Rank        int16  `json:"rank"`
}

func (q *Queries) CreateWatchlistCoins(ctx context.Context, arg CreateWatchlistCoinsParams) (WatchlistCoin, error) {
	row := q.db.QueryRow(ctx, createWatchlistCoins,
		arg.WatchlistID,
		arg.Name,
		arg.Symbol,
		arg.Rank,
	)
	var i WatchlistCoin
	err := row.Scan(
		&i.ID,
		&i.WatchlistID,
		&i.Name,
		&i.Symbol,
		&i.Rank,
	)
	return i, err
}

const deleteWatchlistCoin = `-- name: DeleteWatchlistCoin :exec
DELETE FROM watchlist_coins
WHERE id = $1
`

func (q *Queries) DeleteWatchlistCoin(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteWatchlistCoin, id)
	return err
}

const getWatchlistCoin = `-- name: GetWatchlistCoin :one
SELECT id, watchlist_id, name, symbol, rank FROM watchlist_coins
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetWatchlistCoin(ctx context.Context, id int64) (WatchlistCoin, error) {
	row := q.db.QueryRow(ctx, getWatchlistCoin, id)
	var i WatchlistCoin
	err := row.Scan(
		&i.ID,
		&i.WatchlistID,
		&i.Name,
		&i.Symbol,
		&i.Rank,
	)
	return i, err
}

const listWatchlistsCoins = `-- name: ListWatchlistsCoins :many
SELECT id, watchlist_id, name, symbol, rank FROM watchlist_coins
WHERE watchlist_id = $1
`

func (q *Queries) ListWatchlistsCoins(ctx context.Context, watchlistID int64) ([]WatchlistCoin, error) {
	rows, err := q.db.Query(ctx, listWatchlistsCoins, watchlistID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []WatchlistCoin{}
	for rows.Next() {
		var i WatchlistCoin
		if err := rows.Scan(
			&i.ID,
			&i.WatchlistID,
			&i.Name,
			&i.Symbol,
			&i.Rank,
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
