package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nico-mayer/go-api/config"
	"github.com/nico-mayer/go-api/controllers"
	"github.com/nico-mayer/go-api/db"
	"github.com/nico-mayer/go-api/models"
)

func main() {

	models.People = append(models.People, models.Person{ID: "1", FirstName: "John", LastName: "Doe", Age: 30})
	models.People = append(models.People, models.Person{ID: "2", FirstName: "Jane", LastName: "Doe", Age: 28})

	err := db.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	handleRequests()
}

func handleRequests() {
	http.HandleFunc("/people", controllers.GetPeople)
	http.HandleFunc("/people/create", controllers.CreatePerson)

	fmt.Println("Server listening on PORT: " + config.PORT)
	log.Fatal(http.ListenAndServe(":"+config.PORT, nil))
}
