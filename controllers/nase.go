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

	var nase *models.Nase

	if err := json.NewDecoder(req.Body).Decode(&nase); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	if err := models.ValidateNase(nase); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	err := models.InsertNase(nase)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusCreated)
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(map[string]string{"message": "Nase created successfully"})
}
