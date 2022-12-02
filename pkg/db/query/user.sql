-- name: CreateUser :one
INSERT INTO users (
    email, name, last_name, password
) VALUES (
             $1, $2, $3, $4
         )
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: UpdateUser :one
UPDATE users
set name = $2,
    last_name = $3,
    password = $4
WHERE id = $1
RETURNING *;


-- name: CreateWatchlist :one
INSERT INTO watchlists (
    name, account_id
) VALUES (
             $1, $2
         )
RETURNING *;

-- name: GetWatchlist :one
SELECT * FROM watchlists
WHERE id = $1 LIMIT 1;

-- name: ListWatchlists :many
SELECT * FROM watchlists
WHERE account_id = $1;

-- name: DeleteWatchlist :exec
DELETE FROM watchlists
WHERE id = $1;

-- name: UpdateWatchlist :one
UPDATE watchlists
set name = $2
WHERE id = $1
RETURNING *;


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


-- name: CreateCoin :one
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
WHERE account_id = $1;

-- name: ListTransactionsByAccountByCoin :many
SELECT * FROM transactions
WHERE account_id = $1 AND coin_id = $2;

-- name: DeleteTransaction :exec
DELETE FROM transactions
WHERE id = $1;


-- name: CreateFootball :one
INSERT INTO football (
    account_id, team, league, country
) VALUES (
             $1, $2, $3, $4
         )
RETURNING *;

-- name: GetFootball :one
SELECT * FROM football
WHERE account_id = $1 LIMIT 1;


-- name: UpdateFootball :one
UPDATE football
set team = $2,
    league = $3,
    country = $4
WHERE account_id = $1
RETURNING *;