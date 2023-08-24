package crypto

import (
	db "github.com/jakub/aioportal/server/db/sqlc"
	"github.com/jakub/aioportal/server/util"
	"net/http"
)

const (
	rapidApiHeaderKey  = "X-RapidAPI-Key"
	rapidApiHeaderHost = "X-RapidAPI-Host"
)

type CryptoProcessor interface {
	Start() error
}

type CryptoScraperProcessor struct {
	client *http.Client
	config util.Config
	store  db.Store
}

func NewCryptoProcessor(client *http.Client, config util.Config, store db.Store) CryptoProcessor {
	return &CryptoScraperProcessor{
		client: client,
		config: config,
		store:  store,
	}
}

func (processor *CryptoScraperProcessor) Start() error {
	coins, err := processor.getCoinsListRequest()
	if err != nil {
		return err
	}

	return processor.processCoins(coins)
}
