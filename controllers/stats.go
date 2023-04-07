package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/nico-mayer/go-api/models"
)

func GetLeaderboard(res http.ResponseWriter, req *http.Request) {
	var leaderboard []models.User
	if req.Method != "GET" {
		http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	users, err := models.GetAllUsers()
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, user := range users {
		count, err := models.CountNasen(user.ID)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		user.Count = count
		leaderboard = append(leaderboard, user)
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(leaderboard)
}
