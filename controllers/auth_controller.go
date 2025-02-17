package controllers

import (
	"encoding/json"
	"net/http"
	"event-ticketing-system-backend/models"
	"event-ticketing-system-backend/utils"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate user (pseudo-code for now)
	// if !validateUser(user) { }

	token, err := utils.GenerateJWT(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
