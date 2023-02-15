// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: transaction.sql

package db

import (
	"context"
	"time"
)

const createTransaction = `-- name: CreateTransaction :one
INSERT INTO transactions (
    account_id, coin_id, coin_name, symbol, type, amount, time_transacted, price_purchased_at, no_of_coins
) VALUES (
             $1, $2, $3, $4, $5, $6, $7, $8, $9
         )
RETURNING id, account_id, coin_id, coin_name, symbol, type, amount, time_transacted, time_created, price_purchased_at, no_of_coins
`

type CreateTransactionParams struct {
	AccountID        int64     `json:"account_id"`
	CoinID           int64     `json:"coin_id"`
	CoinName         string    `json:"coin_name"`
	Symbol           string    `json:"symbol"`
	Type             int32     `json:"type"`
	Amount           int32     `json:"amount"`
	TimeTransacted   time.Time `json:"time_transacted"`
	PricePurchasedAt string    `json:"price_purchased_at"`
	NoOfCoins        string    `json:"no_of_coins"`
}

func (q *Queries) CreateTransaction(ctx context.Context, arg CreateTransactionParams) (Transaction, error) {
	row := q.db.QueryRowContext(ctx, createTransaction,
		arg.AccountID,
		arg.CoinID,
		arg.CoinName,
		arg.Symbol,
		arg.Type,
		arg.Amount,
		arg.TimeTransacted,
		arg.PricePurchasedAt,
		arg.NoOfCoins,
	)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.CoinID,
		&i.CoinName,
		&i.Symbol,
		&i.Type,
		&i.Amount,
		&i.TimeTransacted,
		&i.TimeCreated,
		&i.PricePurchasedAt,
		&i.NoOfCoins,
	)
	return i, err
}

const deleteTransaction = `-- name: DeleteTransaction :exec





DELETE FROM transactions
WHERE id = $1
`

// -- name: ListTransactionsByAccountByCoin :many
// SELECT * FROM transactions
// WHERE account_id = $1 AND coin_id = $2;
// -- name: ListTransactionsByAccountByCoin :many
// SELECT * FROM transactions
// ORDER BY id
// LIMIT $3
// WHERE account_id = $1 AND coin_id = $2;
func (q *Queries) DeleteTransaction(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteTransaction, id)
	return err
}

const getTransaction = `-- name: GetTransaction :one
SELECT id, account_id, coin_id, coin_name, symbol, type, amount, time_transacted, time_created, price_purchased_at, no_of_coins FROM transactions
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTransaction(ctx context.Context, id int64) (Transaction, error) {
	row := q.db.QueryRowContext(ctx, getTransaction, id)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.CoinID,
		&i.CoinName,
		&i.Symbol,
		&i.Type,
		&i.Amount,
		&i.TimeTransacted,
		&i.TimeCreated,
		&i.PricePurchasedAt,
		&i.NoOfCoins,
	)
	return i, err
}

const listTransactionsByAccount = `-- name: ListTransactionsByAccount :many
SELECT id, account_id, coin_id, coin_name, symbol, type, amount, time_transacted, time_created, price_purchased_at, no_of_coins FROM transactions
WHERE account_id = $1
`

func (q *Queries) ListTransactionsByAccount(ctx context.Context, accountID int64) ([]Transaction, error) {
	rows, err := q.db.QueryContext(ctx, listTransactionsByAccount, accountID)
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
			&i.CoinID,
			&i.CoinName,
			&i.Symbol,
			&i.Type,
			&i.Amount,
			&i.TimeTransacted,
			&i.TimeCreated,
			&i.PricePurchasedAt,
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
