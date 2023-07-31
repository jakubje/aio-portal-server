-- name: CreateUser :one
INSERT INTO users (
    email, name, last_name, password, password_changed_at
) VALUES (
             $1, $2, $3, $4, $5
         )
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

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
SET
    email = COALESCE(sqlc.narg(email), email),

    name = COALESCE(sqlc.narg(name), name),

    last_name = COALESCE(sqlc.narg(last_name), last_name),

    password = COALESCE(sqlc.narg(password), password)

WHERE id = sqlc.arg(id)
RETURNING *;
