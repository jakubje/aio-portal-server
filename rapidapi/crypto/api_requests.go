package crypto

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/jakub/aioportal/server/db/sqlc"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	baseUrl = "https://coinranking1.p.rapidapi.com"
)

func (processor *CryptoScraperProcessor) getCoinsListRequest() (data CoinsResponse, err error) {
	requestUrl := fmt.Sprintf("%s/coins", baseUrl)
	params := url.Values{}
	params.Add("tiers[0]", "1")
	params.Add("orderBy", "marketCap")
	params.Add("orderDirection", "desc")
	params.Add("limit", "1000")
	//tier 2 has 1056 coins
	//tier 3 has
	constructedURL := requestUrl + "?" + params.Encode()

	var responseData CoinsResponse
	req, _ := http.NewRequest("GET", constructedURL, nil)

	req.Header.Add(rapidApiHeaderKey, processor.config.XRapidAPIKey)
	req.Header.Add(rapidApiHeaderHost, processor.config.XRapidAPICryptoHost)

	res, err := processor.client.Do(req)
	if err != nil {
		return responseData, fmt.Errorf("failed to complete request: %s", err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return responseData, fmt.Errorf("failed to decode body: %s", err)
	}

	err = json.Unmarshal(body, &responseData)
	if err != nil {
		return responseData, fmt.Errorf("error decoding json: %s", err)
	}

	return responseData, nil
}

func (processor *CryptoScraperProcessor) processCoins(response CoinsResponse) error {

	for _, coin := range response.Data.Coins {
		time.Sleep(2 * time.Second)
		err := processor.processIndividualCoin(coin.UUID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (processor *CryptoScraperProcessor) processIndividualCoin(coinId string) error {

	url := fmt.Sprintf("https://coinranking1.p.rapidapi.com/coin/%s?timePeriod=24h", coinId)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add(rapidApiHeaderKey, processor.config.XRapidAPIKey)
	req.Header.Add(rapidApiHeaderHost, processor.config.XRapidAPICryptoHost)

	res, err := processor.client.Do(req)
	if err != nil {
		return status.Errorf(codes.Internal, "failed to call api: %s", err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("failed to decode body: %s", err)
	}

	var responseData CoinResponse
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		return fmt.Errorf("error decoding json: %s", err)
	}

	coin := responseData.Data.Coin

	var links []string

	for _, link := range coin.Links {
		links = append(links, link.URL)
	}
	//tableEmpty := processor.store.CheckCoins()
	tableEmpty := false

	if tableEmpty {
		arg := db.CreateCoinParams{
			CoinID:            coin.Symbol,
			Name:              coin.Name,
			Price:             coin.Price,
			MarketCap:         coin.MarketCap,
			CirculatingSupply: coin.Supply.Circulating,
			TotalSupply:       coin.Supply.Total,
			MaxSupply:         coin.Supply.Max,
			Rank:              strconv.Itoa(coin.Rank),
			Volume:            coin.Two4HVolume,
			ImageUrl:          coin.IconURL,
			Description:       coin.Description,
			Website:           coin.WebsiteURL,
			SocialMediaLinks:  links,
			UpdatedAt:         time.Now(),
		}

		_, err = processor.store.CreateCoin(context.Background(), arg)
		if err != nil {
			return status.Errorf(codes.Internal, "failed to create coin: %s", err)
		}
	} else {
		log.Info().Msg(fmt.Sprintf("Updating coin %s", coin.Symbol))
		arg := db.UpdateCoinParams{
			Name:              pgtype.Text{String: coin.Name, Valid: true},
			Price:             pgtype.Text{String: coin.Price, Valid: true},
			MarketCap:         pgtype.Text{String: coin.MarketCap, Valid: true},
			CirculatingSupply: pgtype.Text{String: coin.Supply.Circulating, Valid: true},
			TotalSupply:       pgtype.Text{String: coin.Supply.Total, Valid: true},
			MaxSupply:         pgtype.Text{String: coin.Supply.Max, Valid: true},
			Rank:              pgtype.Text{String: strconv.Itoa(coin.Rank), Valid: true},
			Volume:            pgtype.Text{String: coin.Two4HVolume, Valid: true},
			ImageUrl:          pgtype.Text{String: coin.IconURL, Valid: true},
			Description:       pgtype.Text{String: coin.Description, Valid: true},
			Website:           pgtype.Text{String: coin.WebsiteURL, Valid: true},
			SocialMediaLinks:  links,
			UpdatedAt: pgtype.Timestamptz{
				Time:  time.Now(),
				Valid: true,
			},
			CoinID: coin.Symbol,
		}
		_, err := processor.store.UpdateCoin(context.Background(), arg)
		if err != nil {
			return status.Errorf(codes.Internal, "failed to update coin: %s", err)
		}
	}

	return nil
}
