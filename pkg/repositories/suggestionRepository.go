package repositories

import (
	"github.com/turistaremsampa/hackaton-visitesp-backend/pkg/database"
	"github.com/turistaremsampa/hackaton-visitesp-backend/pkg/models"
)

func CreateSuggestion(suggestion *models.Suggestion) error {
	return database.DB.Create(suggestion).Error
}

func GetAllSuggestions() ([]models.Suggestion, error) {
	var suggestions []models.Suggestion
	if err := database.DB.Find(&suggestions).Error; err != nil {
		return nil, err
	}
	return suggestions, nil
}

func GetSuggestionByID(id string) (models.Suggestion, error) {
	var suggestion models.Suggestion
	if err := database.DB.First(&suggestion, id).Error; err != nil {
		return suggestion, err
	}
	return suggestion, nil
}
