package db

import (
	"context"
	"github.com/jakub/aioportal/server/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func CreateRandomCoin(t *testing.T) Coin {

	var socialMediaLinks []string
	for i := 0; i < 10; i++ {
		socialMediaLinks = append(socialMediaLinks, util.RandomString(10))
	}

	arg := CreateCoinParams{
		CoinID:            util.RandomString(3),
		Name:              util.RandomString(6),
		Price:             float64(util.RandomInt()),
		MarketCap:         int64(util.RandomInt()),
		CirculatingSupply: int64(util.RandomInt()),
		TotalSupply:       int64(util.RandomInt()),
		MaxSupply:         int64(util.RandomInt()),
		Rank:              int32(util.RandomInt()),
		Volume:            int64(util.RandomInt()),
		ImageUrl:          util.RandomString(10),
		Description:       util.RandomString(100),
		Website:           util.RandomString(15),
		SocialMediaLinks:  socialMediaLinks,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	coin, err := testStore.CreateCoin(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, coin)

	require.Equal(t, arg.CoinID, coin.CoinID)
	require.Equal(t, arg.Name, coin.Name)
	require.Equal(t, arg.Price, coin.Price)
	require.Equal(t, arg.MarketCap, coin.MarketCap)
	require.Equal(t, arg.CirculatingSupply, coin.CirculatingSupply)
	require.Equal(t, arg.TotalSupply, coin.TotalSupply)
	require.Equal(t, arg.MaxSupply, coin.MaxSupply)
	require.Equal(t, arg.Rank, coin.Rank)
	require.Equal(t, arg.Volume, coin.Volume)
	require.Equal(t, arg.ImageUrl, coin.ImageUrl)
	require.Equal(t, arg.Description, coin.Description)
	require.Equal(t, arg.Website, coin.Website)
	require.Equal(t, arg.SocialMediaLinks, coin.SocialMediaLinks)
	require.WithinDuration(t, arg.CreatedAt, coin.CreatedAt, time.Second)

	return coin
}

func TestCreateRandomCoin(t *testing.T) {
	CreateRandomCoin(t)
}
