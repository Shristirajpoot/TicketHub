package routes

import (
	"net/http"
	"event-ticketing-system-backend/controllers"
)

func Login(w http.ResponseWriter, r *http.Request) {
	controllers.Login(w, r)
}
