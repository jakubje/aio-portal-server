-- name: CreateCoin :one
INSERT INTO coins (
    coin_id, coin_uuid, name, price, market_cap, number_of_markets, number_of_exchanges, approved_supply, circulating_supply, total_supply, max_supply, rank, volume, daily_change, image_url, description, all_time_high, tags, website, social_media_links, updated_at
) VALUES (
             $1, $2, $3, $4, $5, $6, $7, $8 ,$9 ,$10 ,$11 ,$12 ,$13 ,$14, $15, $16, $17, $18, $19, $20, $21
         ) ON CONFLICT (coin_id) DO UPDATE
SET
    coin_uuid = $2,
    name = $3,
    price = $4,
    market_cap = $5,
    number_of_markets = $6,
    number_of_exchanges = $7,
    approved_supply = $8,
    circulating_supply = $9,
    total_supply =  $10,
    max_supply = $11,
    rank = $12,
    volume = $13,
    daily_change = $14,
    image_url = $15,
    description = $16,
    all_time_high = $17,
    tags = $18,
    website = $19,
    social_media_links = $20,
    updated_at = $21
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
