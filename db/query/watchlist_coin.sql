-- name: CreateWatchlistCoins :one
INSERT INTO watchlist_coins (
    watchlist_id, name, symbol, rank
) VALUES (
             $1, $2, $3, $4
         )
RETURNING *;

-- name: GetWatchlistCoin :one
SELECT * FROM watchlist_coins
WHERE id = $1 LIMIT 1;

-- name: ListWatchlistsCoins :many
SELECT * FROM watchlist_coins
WHERE watchlist_id = $1;

-- name: DeleteWatchlistCoin :exec
DELETE FROM watchlist_coins
WHERE id = $1;