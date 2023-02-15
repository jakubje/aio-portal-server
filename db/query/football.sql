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