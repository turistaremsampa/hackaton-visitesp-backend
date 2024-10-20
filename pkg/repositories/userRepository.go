package repositories

import (
	"github.com/turistaremsampa/hackaton-visitesp-backend/pkg/database"
	"github.com/turistaremsampa/hackaton-visitesp-backend/pkg/models"
)

func CreateUser(user *models.User) error {
	return database.DB.Create(user).Error
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := database.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
