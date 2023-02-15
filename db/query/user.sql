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

-- -- name: UpdateUser :one
-- UPDATE users
-- set 
--     email = CASE WHEN $2 IS NOT NULL THEN $2 ELSE email END,
--     name = CASE WHEN $3 IS NOT NULL THEN $3 ELSE name END,
--     last_name = CASE WHEN $4 IS NOT NULL THEN $4 ELSE last_name END,
--     password = CASE WHEN $5 IS NOT NULL THEN $5 ELSE password END
-- WHERE id = $1
-- RETURNING *;

-- name: UpdateUser :one
UPDATE users
set 
    email = COALESCE($2, email), 
    name = COALESCE($3, name), 
    last_name = COALESCE($4, last_name), 
    password = COALESCE($5, password)
WHERE id = $1
RETURNING *;
