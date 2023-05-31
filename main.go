package main

import (
	"database/sql"
	"log"
	"server/api"
	db "server/db/sqlc"
	"server/util"

	_ "github.com/lib/pq"
)

func main() {
	// router := routes.UsersResource{}.Routes()
	// fmt.Println("Starting server on the port 3333...")
	// log.Fatal(http.ListenAndServe(":3333", router))
	// store := db.NewStore(conn)

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)

	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

}
