// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: user.sql

package db

import (
	"context"
	"time"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    email, name, last_name, password, password_changed_at
) VALUES (
             $1, $2, $3, $4, $5
         )
RETURNING id, email, name, last_name, password, password_changed_at, created_at
`

type CreateUserParams struct {
	Email             string    `json:"email"`
	Name              string    `json:"name"`
	LastName          string    `json:"last_name"`
	Password          string    `json:"password"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Email,
		arg.Name,
		arg.LastName,
		arg.Password,
		arg.PasswordChangedAt,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Name,
		&i.LastName,
		&i.Password,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, email, name, last_name, password, password_changed_at, created_at FROM users
WHERE email = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Name,
		&i.LastName,
		&i.Password,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, email, name, last_name, password, password_changed_at, created_at FROM users
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListUsersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.Name,
			&i.LastName,
			&i.Password,
			&i.PasswordChangedAt,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
set 
    email = COALESCE($2, email), 
    name = COALESCE($3, name), 
    last_name = COALESCE($4, last_name), 
    password = COALESCE($5, password),
    password_changed_at = COALESCE($6, password_changed_at)

WHERE id = $1
RETURNING id, email, name, last_name, password, password_changed_at, created_at
`

type UpdateUserParams struct {
	ID                int64     `json:"id"`
	Email             string    `json:"email"`
	Name              string    `json:"name"`
	LastName          string    `json:"last_name"`
	Password          string    `json:"password"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.ID,
		arg.Email,
		arg.Name,
		arg.LastName,
		arg.Password,
		arg.PasswordChangedAt,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Name,
		&i.LastName,
		&i.Password,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}
