package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/nico-mayer/go-api/models"
)

func CreateUser(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newUser models.User

	if err := json.NewDecoder(req.Body).Decode(&newUser); err != nil {
		http.Error(res, "Bad request", http.StatusBadRequest)
		return
	}

	if err := models.ValidateUser(&newUser); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	err := models.InsertUser(newUser)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusCreated)
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(map[string]string{"message": "User created successfully"})
}
