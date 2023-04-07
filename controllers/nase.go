package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/nico-mayer/go-api/models"
)

type createRequest struct {
	UserIDs  []string `json:"users"`
	AuthorID string   `json:"authorid"`
	Reason   string   `json:"reason"`
}

func CreateNase(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var reqBody createRequest

	if err := json.NewDecoder(req.Body).Decode(&reqBody); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	if len(reqBody.UserIDs) == 0 {
		http.Error(res, "Missing or invalid users", http.StatusBadRequest)
		return
	}

	for _, id := range reqBody.UserIDs {
		newUUID, err := uuid.NewRandom()
		if err != nil {
			fmt.Println("Error generating UUID:", err)
			continue
		}

		newNase := models.Nase{
			ID:       newUUID,
			UserID:   id,
			AuthorID: reqBody.AuthorID,
			Reason:   reqBody.Reason,
		}

		go models.InsertNase(&newNase)
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(map[string]string{"message": "Nasen created successfully"})
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
