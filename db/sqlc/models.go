// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"time"

	"github.com/google/uuid"
)

type Football struct {
	ID        int64  `json:"id"`
	AccountID int64  `json:"account_id"`
	Team      string `json:"team"`
	League    string `json:"league"`
	Country   string `json:"country"`
}

type Portfolio struct {
	ID         int64  `json:"id"`
	AccountID  int64  `json:"account_id"`
	Name       string `json:"name"`
	Holdings   int32  `json:"holdings"`
	Change24h  int32  `json:"change_24h"`
	ProfitLoss int32  `json:"profit_loss"`
}

type Transaction struct {
	ID             uuid.UUID `json:"id"`
	AccountID      int64     `json:"account_id"`
	PortfolioID    int64     `json:"portfolio_id"`
	Type           int32     `json:"type"`
	Symbol         string    `json:"symbol"`
	PricePerCoin   float64   `json:"price_per_coin"`
	Quantity       float64   `json:"quantity"`
	TimeTransacted time.Time `json:"time_transacted"`
	TimeCreated    time.Time `json:"time_created"`
}

type User struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	LastName  string    `json:"last_name"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type Watchlist struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	AccountID int64  `json:"account_id"`
}

type WatchlistCoin struct {
	ID          int64  `json:"id"`
	WatchlistID int64  `json:"watchlist_id"`
	Name        string `json:"name"`
	Symbol      string `json:"symbol"`
	Rank        int16  `json:"rank"`
}
