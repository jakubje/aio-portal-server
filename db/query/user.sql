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
set 
    email = COALESCE($2, email), 
    name = COALESCE($3, name), 
    last_name = COALESCE($4, last_name), 
    password = COALESCE($5, password)
WHERE id = $1
RETURNING *;
