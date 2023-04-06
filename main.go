package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/nico-mayer/go-api/controllers"
	"github.com/nico-mayer/go-api/models"
)

func main() {
	models.People = append(models.People, models.Person{ID: "1", FirstName: "John", LastName: "Doe", Age: 30})
	models.People = append(models.People, models.Person{ID: "2", FirstName: "Jane", LastName: "Doe", Age: 28})

	var port string
	err := godotenv.Load()
	if err != nil {
		port = ":8080"
	}
	port = ":" + os.Getenv("PORT")

	handleRequests(port)
}

func handleRequests(port string) {
	http.HandleFunc("/people", controllers.GetPeople)
	http.HandleFunc("/people/create", controllers.CreatePerson)
	fmt.Println("Server listening on PORT: ", port)

	log.Fatal(http.ListenAndServe(port, nil))
}
