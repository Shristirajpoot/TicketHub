package main

import (
	"fmt"
	"log"
	"net/http"
	"event-ticketing-system-backend/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Setup routes
	r.HandleFunc("/api/events", routes.GetEvents).Methods("GET")
	r.HandleFunc("/api/events/{id}", routes.GetEvent).Methods("GET")
	r.HandleFunc("/api/events", routes.CreateEvent).Methods("POST")
	r.HandleFunc("/api/auth/login", routes.Login).Methods("POST")

	// Start server
	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
