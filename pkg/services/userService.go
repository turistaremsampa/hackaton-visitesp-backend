package services

import (
	"github.com/turistaremsampa/hackaton-visitesp-backend/pkg/models"
	"github.com/turistaremsampa/hackaton-visitesp-backend/pkg/repositories"
)

func CreateUser(user *models.User) error {
	return repositories.CreateUser(user)
}

func GetAllUsers() ([]models.User, error) {
	return repositories.GetAllUsers()
}
