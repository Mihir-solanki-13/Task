package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"TASK/internal/models"
)

var (
	DB *gorm.DB
)

func InitDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Auto Migrate the schema
	DB.AutoMigrate(&models.Task{})
}
