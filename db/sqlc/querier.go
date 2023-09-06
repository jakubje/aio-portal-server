// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	AddWatchlistCoin(ctx context.Context, arg AddWatchlistCoinParams) (WatchlistCoin, error)
	CreateCoin(ctx context.Context, arg CreateCoinParams) (Coin, error)
	CreateFootball(ctx context.Context, arg CreateFootballParams) (Football, error)
	CreatePortfolio(ctx context.Context, arg CreatePortfolioParams) (Portfolio, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	CreateTransaction(ctx context.Context, arg CreateTransactionParams) (Transaction, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	CreateVerifyEmail(ctx context.Context, arg CreateVerifyEmailParams) (VerifyEmail, error)
	CreateWatchlist(ctx context.Context, arg CreateWatchlistParams) (Watchlist, error)
	DeletePortfolio(ctx context.Context, arg DeletePortfolioParams) error
	DeleteTransaction(ctx context.Context, arg DeleteTransactionParams) error
	DeleteUser(ctx context.Context, id int64) error
	DeleteWatchlist(ctx context.Context, arg DeleteWatchlistParams) error
	GetCoin(ctx context.Context, coinID string) (Coin, error)
	GetFootball(ctx context.Context, accountID int64) (Football, error)
	GetPortfolio(ctx context.Context, arg GetPortfolioParams) (Portfolio, error)
	GetRollUpByCoinByPortfolio(ctx context.Context, arg GetRollUpByCoinByPortfolioParams) ([]GetRollUpByCoinByPortfolioRow, error)
	GetSession(ctx context.Context, id uuid.UUID) (Session, error)
	GetTransaction(ctx context.Context, arg GetTransactionParams) (Transaction, error)
	GetUser(ctx context.Context, email string) (User, error)
	GetWatchlist(ctx context.Context, arg GetWatchlistParams) (Watchlist, error)
	ListCoins(ctx context.Context, arg ListCoinsParams) ([]Coin, error)
	ListPortforlios(ctx context.Context, accountID int64) ([]Portfolio, error)
	ListTransactionsByAccount(ctx context.Context, arg ListTransactionsByAccountParams) ([]Transaction, error)
	ListTransactionsByAccountByCoin(ctx context.Context, arg ListTransactionsByAccountByCoinParams) ([]Transaction, error)
	ListTransactionsByPortfolio(ctx context.Context, arg ListTransactionsByPortfolioParams) ([]Transaction, error)
	ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error)
	ListWatchlistCoins(ctx context.Context, arg ListWatchlistCoinsParams) ([]Coin, error)
	ListWatchlists(ctx context.Context, accountID int64) ([]Watchlist, error)
	RemoveWatchlistCoin(ctx context.Context, arg RemoveWatchlistCoinParams) error
	UpdateFootball(ctx context.Context, arg UpdateFootballParams) (Football, error)
	UpdatePortfolio(ctx context.Context, arg UpdatePortfolioParams) (Portfolio, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
	UpdateVerifyEmail(ctx context.Context, arg UpdateVerifyEmailParams) (VerifyEmail, error)
	UpdateWatchlist(ctx context.Context, arg UpdateWatchlistParams) (Watchlist, error)
}

var _ Querier = (*Queries)(nil)
