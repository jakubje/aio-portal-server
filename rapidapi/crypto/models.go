package crypto

type CoinsResponse struct {
	Status string `json:"status"`
	Data   struct {
		Stats struct {
			Total          int    `json:"total"`
			TotalCoins     int    `json:"totalCoins"`
			TotalMarkets   int    `json:"totalMarkets"`
			TotalExchanges int    `json:"totalExchanges"`
			TotalMarketCap string `json:"totalMarketCap"`
			Total24HVolume string `json:"total24hVolume"`
		} `json:"stats"`
		Coins []struct {
			UUID           string   `json:"uuid"`
			Symbol         string   `json:"symbol"`
			Name           string   `json:"name"`
			Color          string   `json:"color"`
			IconURL        string   `json:"iconUrl"`
			MarketCap      string   `json:"marketCap"`
			Price          string   `json:"price"`
			ListedAt       int      `json:"listedAt"`
			Tier           int      `json:"tier"`
			Change         string   `json:"change"`
			Rank           int      `json:"rank"`
			Sparkline      []string `json:"sparkline"`
			LowVolume      bool     `json:"lowVolume"`
			CoinrankingURL string   `json:"coinrankingUrl"`
			Two4HVolume    string   `json:"24hVolume"`
			BtcPrice       string   `json:"btcPrice"`
		} `json:"coins"`
	} `json:"data"`
}

type CoinResponse struct {
	Status string `json:"status"`
	Data   struct {
		Coin struct {
			UUID        string `json:"uuid"`
			Symbol      string `json:"symbol"`
			Name        string `json:"name"`
			Description string `json:"description"`
			Color       string `json:"color"`
			IconURL     string `json:"iconUrl"`
			WebsiteURL  string `json:"websiteUrl"`
			Links       []struct {
				Name string `json:"name"`
				Type string `json:"type"`
				URL  string `json:"url"`
			} `json:"links"`
			Supply struct {
				Confirmed   bool   `json:"confirmed"`
				SupplyAt    int    `json:"supplyAt"`
				Max         string `json:"max"`
				Total       string `json:"total"`
				Circulating string `json:"circulating"`
			} `json:"supply"`
			NumberOfMarkets       int32    `json:"numberOfMarkets"`
			NumberOfExchanges     int32    `json:"numberOfExchanges"`
			Two4HVolume           string   `json:"24hVolume"`
			MarketCap             string   `json:"marketCap"`
			FullyDilutedMarketCap string   `json:"fullyDilutedMarketCap"`
			Price                 string   `json:"price"`
			BtcPrice              string   `json:"btcPrice"`
			PriceAt               int      `json:"priceAt"`
			Change                string   `json:"change"`
			Rank                  int32    `json:"rank"`
			Sparkline             []string `json:"sparkline"`
			AllTimeHigh           struct {
				Price     string `json:"price"`
				Timestamp int    `json:"timestamp"`
			} `json:"allTimeHigh"`
			CoinrankingURL string   `json:"coinrankingUrl"`
			Tier           int      `json:"tier"`
			LowVolume      bool     `json:"lowVolume"`
			ListedAt       int      `json:"listedAt"`
			HasContent     bool     `json:"hasContent"`
			Notices        any      `json:"notices"`
			Tags           []string `json:"tags"`
		} `json:"coin"`
	} `json:"data"`
}
