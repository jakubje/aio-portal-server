-- name: AddCoin :one
INSERT INTO coin (
    coin_name, coin_symbol, quantity, time_created
) VALUES (
             $1, $2, $3, $4
         )
RETURNING *;

-- name: GetCoin :one
SELECT * FROM coin
WHERE id = $1 LIMIT 1;

-- name: ListCoins :many
SELECT * FROM coin;

-- name: DeleteCoin :exec
DELETE FROM coin
WHERE id = $1;

-- name: UpdateCoin :one
UPDATE coin
set quantity = $2
WHERE id = $1
RETURNING *;