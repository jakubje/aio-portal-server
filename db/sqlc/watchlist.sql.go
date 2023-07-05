// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: watchlist.sql

package db

import (
	"context"
)

const createWatchlist = `-- name: CreateWatchlist :one
INSERT INTO watchlists (
    name, account_id
) VALUES (
             $1, $2
         )
RETURNING id, name, account_id
`

type CreateWatchlistParams struct {
	Name      string `json:"name"`
	AccountID int64  `json:"account_id"`
}

func (q *Queries) CreateWatchlist(ctx context.Context, arg CreateWatchlistParams) (Watchlist, error) {
	row := q.db.QueryRowContext(ctx, createWatchlist, arg.Name, arg.AccountID)
	var i Watchlist
	err := row.Scan(&i.ID, &i.Name, &i.AccountID)
	return i, err
}

const deleteWatchlist = `-- name: DeleteWatchlist :exec
DELETE FROM watchlists
WHERE id = $1 and account_id = $2
`

type DeleteWatchlistParams struct {
	ID        int64 `json:"id"`
	AccountID int64 `json:"account_id"`
}

func (q *Queries) DeleteWatchlist(ctx context.Context, arg DeleteWatchlistParams) error {
	_, err := q.db.ExecContext(ctx, deleteWatchlist, arg.ID, arg.AccountID)
	return err
}

const getWatchlist = `-- name: GetWatchlist :one
SELECT id, name, account_id FROM watchlists
WHERE id = $1 and account_id = $2
LIMIT 1
`

type GetWatchlistParams struct {
	ID        int64 `json:"id"`
	AccountID int64 `json:"account_id"`
}

func (q *Queries) GetWatchlist(ctx context.Context, arg GetWatchlistParams) (Watchlist, error) {
	row := q.db.QueryRowContext(ctx, getWatchlist, arg.ID, arg.AccountID)
	var i Watchlist
	err := row.Scan(&i.ID, &i.Name, &i.AccountID)
	return i, err
}

const listWatchlists = `-- name: ListWatchlists :many
SELECT id, name, account_id FROM watchlists
WHERE account_id = $1
`

func (q *Queries) ListWatchlists(ctx context.Context, accountID int64) ([]Watchlist, error) {
	rows, err := q.db.QueryContext(ctx, listWatchlists, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Watchlist{}
	for rows.Next() {
		var i Watchlist
		if err := rows.Scan(&i.ID, &i.Name, &i.AccountID); err != nil {
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

const updateWatchlist = `-- name: UpdateWatchlist :one
UPDATE watchlists
set name = $2
WHERE id = $1 and account_id = $3
RETURNING id, name, account_id
`

type UpdateWatchlistParams struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	AccountID int64  `json:"account_id"`
}

func (q *Queries) UpdateWatchlist(ctx context.Context, arg UpdateWatchlistParams) (Watchlist, error) {
	row := q.db.QueryRowContext(ctx, updateWatchlist, arg.ID, arg.Name, arg.AccountID)
	var i Watchlist
	err := row.Scan(&i.ID, &i.Name, &i.AccountID)
	return i, err
}
