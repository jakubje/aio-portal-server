-- name: CreateTransaction :one
INSERT INTO transactions (
    account_id, coin_id, coin_name, symbol, type, amount, time_transacted, price_purchased_at, no_of_coins
) VALUES (
             $1, $2, $3, $4, $5, $6, $7, $8, $9
         )
RETURNING *;

-- name: GetTransaction :one
SELECT * FROM transactions
WHERE id = $1 LIMIT 1;

-- name: ListTransactionsByAccount :many
SELECT * FROM transactions
WHERE account_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: ListTransactionsByAccountByCoin :many
SELECT * FROM transactions
WHERE account_id = $1 AND coin_id = $2
ORDER BY id
LIMIT $3
OFFSET $4;

-- name: DeleteTransaction :exec
DELETE FROM transactions
WHERE id = $1;