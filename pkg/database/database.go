package database

import (
	"log"

	"github.com/turistaremsampa/hackaton-visitesp-backend/pkg/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("tourism_app.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!")
	}

	// Run auto migrations
	DB.AutoMigrate(&models.User{}, &models.Feedback{}, &models.Suggestion{})
}
