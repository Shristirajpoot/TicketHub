package controllers

import (
	"encoding/json"
	"net/http"
	"event-ticketing-system-backend/models"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func GetEvents(w http.ResponseWriter, r *http.Request) {
	var events []models.Event
	if err := db.Find(&events).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(events)
}

func GetEvent(id string, w http.ResponseWriter, r *http.Request) {
	var event models.Event
	if err := db.First(&event, id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(event)
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	var event models.Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := db.Create(&event).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(event)
}
