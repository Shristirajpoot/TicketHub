package routes

import (
	"encoding/json"
	"net/http"
	"event-ticketing-system-backend/controllers"
	"github.com/gorilla/mux"
)

func GetEvents(w http.ResponseWriter, r *http.Request) {
	controllers.GetEvents(w, r)
}

func GetEvent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	controllers.GetEvent(params["id"], w, r)
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	controllers.CreateEvent(w, r)
}
