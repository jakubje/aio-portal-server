package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jakub/aioportal/server/util"
	"log"
	"os"
	"testing"
)

var testStore Store

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	testStore = NewStore(connPool)
	os.Exit(m.Run())
}
