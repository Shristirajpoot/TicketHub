package utils

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func InitDB() *gorm.DB {
	var err error
	db, err = gorm.Open("sqlite3", "events.db")
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
	}
	return db
}
