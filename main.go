package main

import (
	"database/sql"
	"log"
	"server/api"
	db "server/db/sqlc"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:root@localhost:5432/aio_portal?sslmode=disable"
	serverAddress = "0.0.0.0:8888"
)

func main() {
	// router := routes.UsersResource{}.Routes()
	// fmt.Println("Starting server on the port 3333...")
	// log.Fatal(http.ListenAndServe(":3333", router))
	// store := db.NewStore(conn)

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
 
	store := db.NewStore(conn)
	server := api.NewServer(store)


	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

}
