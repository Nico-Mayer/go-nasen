package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/nico-mayer/go-api/config"
	"github.com/nico-mayer/go-api/controllers"
	"github.com/nico-mayer/go-api/models"
)

func main() {
	connStr := "user=" + config.PGUSER + " password=" + config.PGPASSWORD + " dbname=" + config.PGDATABASE + " host=" + config.PGHOST + ":" + config.PGPORT + " sslmode=disable"
	fmt.Println(connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Successfully connected to PostgreSQL database!")

	models.People = append(models.People, models.Person{ID: "1", FirstName: "John", LastName: "Doe", Age: 30})
	models.People = append(models.People, models.Person{ID: "2", FirstName: "Jane", LastName: "Doe", Age: 28})

	handleRequests()
}

func handleRequests() {
	http.HandleFunc("/people", controllers.GetPeople)
	http.HandleFunc("/people/create", controllers.CreatePerson)

	fmt.Println("Server listening on PORT: " + config.PORT)
	log.Fatal(http.ListenAndServe(":"+config.PORT, nil))
}
