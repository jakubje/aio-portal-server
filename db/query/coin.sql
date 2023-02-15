-- name: AddCoin :one
INSERT INTO coins (
    portfolio_id, coin_name, coin_symbol, amount, no_of_coins
) VALUES (
             $1, $2, $3, $4, $5
         )
RETURNING *;

-- name: GetCoin :one
SELECT * FROM coins
WHERE id = $1 LIMIT 1;

-- name: ListCoins :many
SELECT * FROM coins
WHERE portfolio_id = $1;

-- name: DeleteCoin :exec
DELETE FROM coins
WHERE id = $1;

-- name: UpdateCoin :one
UPDATE coins
set amount = $2,
    no_of_coins = $3
WHERE id = $1
RETURNING *;