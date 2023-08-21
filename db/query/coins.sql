-- name: CreateCoin :one
INSERT INTO coins (
    coin_id, name, price, market_cap, circulation_supply, total_supply, max_supply, rank, volume, image_url, description, website, social_media_links, created_at, updated_at
) VALUES (
             $1, $2, $3, $4, $5, $6, $7, $8 ,$9 ,$10 ,$11 ,$12 ,$13 ,$14 ,$15
         )
RETURNING *;

-- name: GetCoin :one
SELECT * FROM coins
WHERE coin_id = $1
LIMIT 1;

-- name: ListCoins :many
SELECT * FROM coins
ORDER BY rank
LIMIT $1
OFFSET $2;
