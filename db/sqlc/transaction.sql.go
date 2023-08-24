// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: transaction.sql

package db

import (
	"context"
	"github.com/google/uuid"
	"time"

)

const createTransaction = `-- name: CreateTransaction :one
INSERT INTO transactions (
    account_id, portfolio_id, symbol, type, price_per_coin, quantity, time_transacted, time_created
) VALUES (
             $1, $2, $3, $4, $5, $6, $7, $8
         )
RETURNING id, account_id, portfolio_id, type, symbol, price_per_coin, quantity, time_transacted, time_created
`

type CreateTransactionParams struct {
	AccountID      int64     `json:"account_id"`
	PortfolioID    int64     `json:"portfolio_id"`
	Symbol         string    `json:"symbol"`
	Type           int32     `json:"type"`
	PricePerCoin   float64   `json:"price_per_coin"`
	Quantity       float64   `json:"quantity"`
	TimeTransacted time.Time `json:"time_transacted"`
	TimeCreated    time.Time `json:"time_created"`
}

func (q *Queries) CreateTransaction(ctx context.Context, arg CreateTransactionParams) (Transaction, error) {
	row := q.db.QueryRow(ctx, createTransaction,
		arg.AccountID,
		arg.PortfolioID,
		arg.Symbol,
		arg.Type,
		arg.PricePerCoin,
		arg.Quantity,
		arg.TimeTransacted,
		arg.TimeCreated,
	)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.PortfolioID,
		&i.Type,
		&i.Symbol,
		&i.PricePerCoin,
		&i.Quantity,
		&i.TimeTransacted,
		&i.TimeCreated,
	)
	return i, err
}

const deleteTransaction = `-- name: DeleteTransaction :exec
DELETE FROM transactions
WHERE id = $1 and account_id = $2
`

type DeleteTransactionParams struct {
	ID        uuid.UUID `json:"id"`
	AccountID int64       `json:"account_id"`
}

func (q *Queries) DeleteTransaction(ctx context.Context, arg DeleteTransactionParams) error {
	_, err := q.db.Exec(ctx, deleteTransaction, arg.ID, arg.AccountID)
	return err
}

const getRollUpByCoinByPortfolio = `-- name: GetRollUpByCoinByPortfolio :many
SELECT 
symbol, type, 
CAST (SUM(price_per_coin) AS FLOAT) AS total_cost, 
CAST (SUM(quantity) AS FLOAT) AS total_coins,
CAST (CAST(SUM(price_per_coin) AS FLOAT) *1.0 / CAST (SUM(quantity) AS FLOAT) AS FLOAT) AS price_per_coin
FROM transactions
WHERE portfolio_id = $1 and account_id = $2
GROUP BY symbol, type
`

type GetRollUpByCoinByPortfolioParams struct {
	PortfolioID int64 `json:"portfolio_id"`
	AccountID   int64 `json:"account_id"`
}

type GetRollUpByCoinByPortfolioRow struct {
	Symbol       string  `json:"symbol"`
	Type         int32   `json:"type"`
	TotalCost    float64 `json:"total_cost"`
	TotalCoins   float64 `json:"total_coins"`
	PricePerCoin float64 `json:"price_per_coin"`
}

func (q *Queries) GetRollUpByCoinByPortfolio(ctx context.Context, arg GetRollUpByCoinByPortfolioParams) ([]GetRollUpByCoinByPortfolioRow, error) {
	rows, err := q.db.Query(ctx, getRollUpByCoinByPortfolio, arg.PortfolioID, arg.AccountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetRollUpByCoinByPortfolioRow{}
	for rows.Next() {
		var i GetRollUpByCoinByPortfolioRow
		if err := rows.Scan(
			&i.Symbol,
			&i.Type,
			&i.TotalCost,
			&i.TotalCoins,
			&i.PricePerCoin,
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

const getTransaction = `-- name: GetTransaction :one
SELECT id, account_id, portfolio_id, type, symbol, price_per_coin, quantity, time_transacted, time_created FROM transactions
WHERE id = $1 and account_id = $2
LIMIT 1
`

type GetTransactionParams struct {
	ID        uuid.UUID `json:"id"`
	AccountID int64       `json:"account_id"`
}

func (q *Queries) GetTransaction(ctx context.Context, arg GetTransactionParams) (Transaction, error) {
	row := q.db.QueryRow(ctx, getTransaction, arg.ID, arg.AccountID)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.PortfolioID,
		&i.Type,
		&i.Symbol,
		&i.PricePerCoin,
		&i.Quantity,
		&i.TimeTransacted,
		&i.TimeCreated,
	)
	return i, err
}

const listTransactionsByAccount = `-- name: ListTransactionsByAccount :many
SELECT id, account_id, portfolio_id, type, symbol, price_per_coin, quantity, time_transacted, time_created FROM transactions
WHERE account_id = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListTransactionsByAccountParams struct {
	AccountID int64 `json:"account_id"`
	Limit     int32 `json:"limit"`
	Offset    int32 `json:"offset"`
}

func (q *Queries) ListTransactionsByAccount(ctx context.Context, arg ListTransactionsByAccountParams) ([]Transaction, error) {
	rows, err := q.db.Query(ctx, listTransactionsByAccount, arg.AccountID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Transaction{}
	for rows.Next() {
		var i Transaction
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.PortfolioID,
			&i.Type,
			&i.Symbol,
			&i.PricePerCoin,
			&i.Quantity,
			&i.TimeTransacted,
			&i.TimeCreated,
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

const listTransactionsByAccountByCoin = `-- name: ListTransactionsByAccountByCoin :many
SELECT id, account_id, portfolio_id, type, symbol, price_per_coin, quantity, time_transacted, time_created FROM transactions
WHERE symbol = $1 and account_id = $2
ORDER BY id
LIMIT $3
OFFSET $4
`

type ListTransactionsByAccountByCoinParams struct {
	Symbol    string `json:"symbol"`
	AccountID int64  `json:"account_id"`
	Limit     int32  `json:"limit"`
	Offset    int32  `json:"offset"`
}

func (q *Queries) ListTransactionsByAccountByCoin(ctx context.Context, arg ListTransactionsByAccountByCoinParams) ([]Transaction, error) {
	rows, err := q.db.Query(ctx, listTransactionsByAccountByCoin,
		arg.Symbol,
		arg.AccountID,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Transaction{}
	for rows.Next() {
		var i Transaction
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.PortfolioID,
			&i.Type,
			&i.Symbol,
			&i.PricePerCoin,
			&i.Quantity,
			&i.TimeTransacted,
			&i.TimeCreated,
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

const listTransactionsByPortfolio = `-- name: ListTransactionsByPortfolio :many
SELECT id, account_id, portfolio_id, type, symbol, price_per_coin, quantity, time_transacted, time_created FROM transactions
WHERE portfolio_id = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListTransactionsByPortfolioParams struct {
	PortfolioID int64 `json:"portfolio_id"`
	Limit       int32 `json:"limit"`
	Offset      int32 `json:"offset"`
}

func (q *Queries) ListTransactionsByPortfolio(ctx context.Context, arg ListTransactionsByPortfolioParams) ([]Transaction, error) {
	rows, err := q.db.Query(ctx, listTransactionsByPortfolio, arg.PortfolioID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Transaction{}
	for rows.Next() {
		var i Transaction
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.PortfolioID,
			&i.Type,
			&i.Symbol,
			&i.PricePerCoin,
			&i.Quantity,
			&i.TimeTransacted,
			&i.TimeCreated,
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
