// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: session.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createSession = `-- name: CreateSession :one
INSERT INTO sessions (
    id,
    account_id,
    email,
     refresh_token,
     user_agent,
     client_ip,
     is_blocked,
     expires_at
) VALUES (
             $1, $2, $3, $4, $5, $6, $7, $8
         )
RETURNING id, email, account_id, refresh_token, user_agent, client_ip, is_blocked, expires_at
`

type CreateSessionParams struct {
	ID           uuid.UUID `json:"id"`
	AccountID    int64     `json:"account_id"`
	Email        string    `json:"email"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiresAt    time.Time `json:"expires_at"`
}

func (q *Queries) CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error) {
	row := q.db.QueryRow(ctx, createSession,
		arg.ID,
		arg.AccountID,
		arg.Email,
		arg.RefreshToken,
		arg.UserAgent,
		arg.ClientIp,
		arg.IsBlocked,
		arg.ExpiresAt,
	)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.AccountID,
		&i.RefreshToken,
		&i.UserAgent,
		&i.ClientIp,
		&i.IsBlocked,
		&i.ExpiresAt,
	)
	return i, err
}

const getSession = `-- name: GetSession :one
SELECT id, email, account_id, refresh_token, user_agent, client_ip, is_blocked, expires_at FROM sessions
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetSession(ctx context.Context, id uuid.UUID) (Session, error) {
	row := q.db.QueryRow(ctx, getSession, id)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.AccountID,
		&i.RefreshToken,
		&i.UserAgent,
		&i.ClientIp,
		&i.IsBlocked,
		&i.ExpiresAt,
	)
	return i, err
}
