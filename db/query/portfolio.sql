-- name: CreatePortfolio :one
INSERT INTO portfolios (
    name, account_id
) VALUES (
             $1, $2
         )
RETURNING *;

-- name: GetPortfolio :one
SELECT * FROM portfolios
WHERE id = $1 LIMIT 1;

-- name: ListPortforlios :many
SELECT * FROM portfolios
WHERE account_id = $1;

-- name: DeletePortfolio :exec
DELETE FROM portfolios
WHERE id = $1;

-- name: UpdatePortfolio :one
UPDATE portfolios
set name = $2
WHERE id = $1
RETURNING *;