package models

import (
	"time"
)

type Event struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	Date           time.Time `json:"date"`
	Location       string    `json:"location"`
	TotalTickets   int       `json:"total_tickets"`
	AvailableTickets int     `json:"available_tickets"`
}
