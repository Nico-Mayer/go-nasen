package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/nico-mayer/go-api/models"
)

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	// Check if method is post
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newPerson models.Person

	// Decode the request body into the newPerson struct
	if err := json.NewDecoder(r.Body).Decode(&newPerson); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Validate the person
	if err := models.ValidatePerson(&newPerson); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	models.People = append(models.People, newPerson)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Person created successfully"})
}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	people := models.GetAllPeople()
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}
