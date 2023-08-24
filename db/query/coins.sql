-- name: CreateCoin :one
INSERT INTO coins (
    coin_id, name, price, market_cap, circulating_supply, total_supply, max_supply, rank, volume, image_url, description, website, social_media_links, updated_at
) VALUES (
             $1, $2, $3, $4, $5, $6, $7, $8 ,$9 ,$10 ,$11 ,$12 ,$13 ,$14
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

-- name: UpdateCoin :one
UPDATE coins
SET
    name = COALESCE(sqlc.narg(name), name),
    price = COALESCE(sqlc.narg(price), price),
    market_cap = COALESCE(sqlc.narg(market_cap), market_cap),
    circulating_supply = COALESCE(sqlc.narg(circulating_supply), circulating_supply),
    total_supply = COALESCE(sqlc.narg(total_supply), total_supply),
    max_supply = COALESCE(sqlc.narg(max_supply), max_supply),
    rank = COALESCE(sqlc.narg(rank), rank),
    volume = COALESCE(sqlc.narg(volume), volume),
    image_url = COALESCE(sqlc.narg(image_url), image_url),
    description = COALESCE(sqlc.narg(description), description),
    website = COALESCE(sqlc.narg(website), website),
    social_media_links = COALESCE(sqlc.narg(social_media_links), social_media_links),
    updated_at = COALESCE(sqlc.narg(updated_at), updated_at)

WHERE coin_id = sqlc.arg(coin_id)
RETURNING *;
