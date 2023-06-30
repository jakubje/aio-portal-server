-- name: CreateWatchlist :one
INSERT INTO watchlists (
    name, account_id
) VALUES (
             $1, $2
         )
RETURNING *;

-- name: GetWatchlist :one
SELECT * FROM watchlists
WHERE id = $1 and account_id = $2
LIMIT 1;

-- name: ListWatchlists :many
SELECT * FROM watchlists
WHERE account_id = $1;

-- name: DeleteWatchlist :exec
DELETE FROM watchlists
WHERE id = $1 and account_id = $2;

-- name: UpdateWatchlist :one
UPDATE watchlists
set name = $2
WHERE id = $1 and account_id = $3
RETURNING *;