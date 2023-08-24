-- name: AddWatchlistCoin :one
INSERT INTO watchlist_coins (
    watchlist_id, coin_id
) VALUES (
             $1, $2
         )
RETURNING *;

-- name: RemoveWatchlistCoin :exec
DELETE FROM watchlist_coins
WHERE watchlist_id = $1 and coin_id = $2;

-- name: ListWatchlistCoins :many
SELECT c.*
FROM coins c
INNER JOIN watchlist_coins wc ON c.coin_id = wc.coin_id
INNER JOIN watchlists watchlist ON wc.watchlist_id = watchlist.id
WHERE wc.watchlist_id = $1 AND watchlist.account_id = $2;