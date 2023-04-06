package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nico-mayer/go-api/config"
	"github.com/nico-mayer/go-api/controllers"
	"github.com/nico-mayer/go-api/db"
)

func main() {
	err := db.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	handleRequests()
}

func handleRequests() {
	http.HandleFunc("/user/create", controllers.CreateUser)

	fmt.Println("Server listening on PORT: " + config.PORT)
	log.Fatal(http.ListenAndServe(":"+config.PORT, nil))
}
