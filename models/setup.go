package models

import (
	"log"
	"os"
	"time"

	"gorm.io/gorm/logger"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

var newLogger = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	logger.Config{
		SlowThreshold:             time.Second,   // Slow SQL threshold
		LogLevel:                  logger.Silent, // Log level
		IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
		Colorful:                  false,         // Disable color
	},
)

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("resume_builder.db"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal(err)
	}
	database.AutoMigrate(&User{}, &Resume{})
	DB = database
}
