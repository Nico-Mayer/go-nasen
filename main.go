package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nico-mayer/go-api/controllers"
	"github.com/nico-mayer/go-api/models"
)

var port = 8080

func main() {
	models.People = append(models.People, models.Person{ID: "1", FirstName: "John", LastName: "Doe", Age: 30})
	models.People = append(models.People, models.Person{ID: "2", FirstName: "Jane", LastName: "Doe", Age: 28})

	handleRequests()
}

func handleRequests() {
	http.HandleFunc("/people", controllers.GetPeople)
	http.HandleFunc("/people/create", controllers.CreatePerson)
	fmt.Println("Server listening on PORT: ", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
