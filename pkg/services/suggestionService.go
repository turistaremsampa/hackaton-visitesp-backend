package services

import (
	"github.com/turistaremsampa/hackaton-visitesp-backend/pkg/models"
	"github.com/turistaremsampa/hackaton-visitesp-backend/pkg/repositories"
)

func CreateSuggestion(suggestion *models.Suggestion) error {
	return repositories.CreateSuggestion(suggestion)
}

func GetAllSuggestions() ([]models.Suggestion, error) {
	return repositories.GetAllSuggestions()
}

func GetSuggestionByID(id string) (models.Suggestion, error) {
	return repositories.GetSuggestionByID(id)
}
