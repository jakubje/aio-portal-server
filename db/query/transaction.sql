-- name: CreateTransaction :one
INSERT INTO transactions (
    account_id, portfolio_id, symbol, type, price_per_coin, amount, quantity, time_transacted, time_created
) VALUES (
             $1, $2, $3, $4, $5, $6, $7, $8, $9
         )
RETURNING *;

-- name: GetTransaction :one
SELECT * FROM transactions
WHERE id = $1 and account_id = $2
LIMIT 1;

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
WHERE symbol = $1 and account_id = $2
ORDER BY id
LIMIT $3
OFFSET $4;

-- name: GetRollUpByCoinByPortfolio :many
SELECT
    t.symbol,
    t.type,
    CAST(SUM(t.amount) AS FLOAT) AS amount,
    CAST(SUM(t.quantity) AS FLOAT) AS total_coins,
    CAST(CAST(SUM(t.price_per_coin) AS FLOAT) * 1.0 / CAST(SUM(t.quantity) AS FLOAT) AS FLOAT) AS price_per_coin,
    CAST(((CAST(c.price AS FLOAT) - CAST(t.price_per_coin AS FLOAT)) / CAST(t.price_per_coin AS FLOAT)) * 100 AS FLOAT) AS profit_loss_percentage
FROM transactions t
         JOIN coins c ON t.symbol = c.coin_id
WHERE t.portfolio_id = $1 AND t.account_id = $2
GROUP BY t.symbol, t.type, c.price, t.price_per_coin;


-- name: DeleteTransaction :exec
DELETE FROM transactions
WHERE id = $1 and account_id = $2 ;