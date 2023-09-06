-- name: CreateCoin :one
INSERT INTO coins (
    coin_id, name, price, market_cap, circulating_supply, total_supply, max_supply, rank, volume, image_url, description, website, social_media_links, updated_at
) VALUES (
             $1, $2, $3, $4, $5, $6, $7, $8 ,$9 ,$10 ,$11 ,$12 ,$13 ,$14
         ) ON CONFLICT (coin_id) DO UPDATE
SET
    name = $2,
    price = $3,
    market_cap = $4,
    circulating_supply = $5,
    total_supply =  $6,
    max_supply = $7,
    rank = $8,
    volume = $9,
    image_url = $10,
    description = $11,
    website = $12,
    social_media_links = $13,
    updated_at = $14
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
