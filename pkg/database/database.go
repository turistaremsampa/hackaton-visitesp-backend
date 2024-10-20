package database

import (
	"fmt"
	"log"

	"github.com/turistaremsampa/hackaton-visitesp-backend/pkg/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	dsn := "root:root@tcp(localhost:3308)/dbVisitesp?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!", DB)
	}

	// Executa migrações automáticas
	err = DB.AutoMigrate(&models.User{}, &models.Feedback{}, &models.Suggestion{})
	if err != nil {
		log.Fatal("Failed to run migrations!")
	}

	fmt.Println("Database connected successfully!")
}
