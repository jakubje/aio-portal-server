package crypto

import (
	"context"
	"encoding/json"
	"fmt"
	db "github.com/jakub/aioportal/server/db/sqlc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"net/http"
	"net/url"
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
	params.Add("limit", "100")
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

	arg := db.CreateCoinParams{
		CoinID:            coin.Symbol,
		CoinUuid:          coin.UUID,
		Name:              coin.Name,
		Price:             coin.Price,
		MarketCap:         coin.MarketCap,
		NumberOfMarkets:   coin.NumberOfMarkets,
		NumberOfExchanges: coin.NumberOfExchanges,
		ApprovedSupply:    coin.Supply.Confirmed,
		CirculatingSupply: coin.Supply.Circulating,
		TotalSupply:       coin.Supply.Total,
		MaxSupply:         coin.Supply.Max,
		Rank:              coin.Rank,
		Volume:            coin.Two4HVolume,
		DailyChange:       coin.Change,
		ImageUrl:          coin.IconURL,
		Description:       coin.Description,
		AllTimeHigh:       coin.AllTimeHigh.Price,
		Tags:              coin.Tags,
		Website:           coin.WebsiteURL,
		SocialMediaLinks:  links,
		UpdatedAt:         time.Now(),
	}

	_, err = processor.store.CreateCoin(context.Background(), arg)
	if err != nil {
		return status.Errorf(codes.Internal, "failed to create/update coin: %s", err)
	}

	return nil
}
