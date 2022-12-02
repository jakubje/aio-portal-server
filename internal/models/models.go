package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID       int64
	Email    string
	Name     string
	LastName string
	Password string
}

type LoginUser struct {
	Email    string
	Password string
}

type UserRegister struct {
	Email    string
	Name     string
	LastName string
	Token    string
}

type AddCoin struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	User         primitive.ObjectID `json:"user,omitempty" bson:"user,omitempty"`
	CoinName     string             `json:"coinName,omitempty" bson:"coinName,omitempty"`
	Symbol       string             `json:"symbol,omitempty" bson:"symbol,omitempty"`
	TotalSpent   int                `json:"totalSpent,omitempty" bson:"totalSpent,omitempty"`
	Quantity     int                `json:"quantity,omitempty" bson:"quantity,omitempty"`
	PricePerCoin int                `json:"pricePerCoin,omitempty" bson:"pricePerCoin,omitempty"`
	Date         time.Time          `json:"date,omitempty" bson:"date,omitempty"`
}

type UpdateCoin struct {
}

//{
//"email",
//"password",
//"portfolio": {
//"favourites": [],
//"assets": {"Ethereum": {"Amount": 2.5}
//}
//},
//"football": {
//"default_league_id": 39,
//'supporting_team': 'Chelsea',
//
//}
//
//}
