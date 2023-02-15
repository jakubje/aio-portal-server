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