package main

import (
	"fmt"
	"log"
	"net/http"
	"server/internal/routes"
)

func main() {

	router := routes.UsersResource{}.Routes()
	fmt.Println("Starting server on the port 3333...")
	log.Fatal(http.ListenAndServe(":3333", router))

	// A good base middleware stack
	//r.Use(middleware.RequestID)
	//r.Use(middleware.RealIP)
	//r.Use(middleware.Logger)
	//r.Use(middleware.Recoverer)
}
