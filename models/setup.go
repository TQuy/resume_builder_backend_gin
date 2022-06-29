package models

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("resume_builder.db"))
	if err != nil {
		log.Fatal(err)
	}
	database.AutoMigrate(&User{}, &Resume{})
	DB = database
}
