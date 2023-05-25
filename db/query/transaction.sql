-- name: CreateTransaction :one
INSERT INTO transactions (
    account_id, portfolio_id, coin_id, symbol, type, price_per_coin, quantity, time_transacted, time_created
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

-- name: ListTransactionsByPortfolio :many
SELECT * FROM transactions
WHERE portfolio_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: ListTransactionsByAccountByCoin :many
SELECT * FROM transactions
WHERE account_id = $1 AND coin_id = $2
ORDER BY id
LIMIT $3
OFFSET $4;

-- name: GetRollUpByCoinByPortfolio :many
SELECT 
symbol, type, 
CAST (SUM(price_per_coin) AS FLOAT) AS total_cost, 
CAST (SUM(quantity) AS FLOAT) AS total_coins 
FROM transactions 
WHERE portfolio_id = $1
GROUP BY symbol, type;

-- name: DeleteTransaction :exec
DELETE FROM transactions
WHERE id = $1;