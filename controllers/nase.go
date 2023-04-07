package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/nico-mayer/go-api/models"
)

func CreateNase(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var nase models.Nase

	if err := json.NewDecoder(req.Body).Decode(&nase); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	if err := models.ValidateNase(&nase); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	err := models.InsertNase(&nase)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(map[string]string{"message": "Nase created successfully"})
}

func GetNasen(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := req.URL.Query().Get("id")

	if id == "" {
		http.Error(res, "Missing or invalid ID", http.StatusBadRequest)
		return
	}

	nasen, err := models.GetNasen(id)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusAccepted)
	json.NewEncoder(res).Encode(nasen)
}

func CountNasen(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := req.URL.Query().Get("id")

	if id == "" {
		http.Error(res, "Missing or invalid ID", http.StatusBadRequest)
		return
	}

	count, err := models.CountNasen(id)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusAccepted)
	json.NewEncoder(res).Encode(map[string]int{"count": count})
}
