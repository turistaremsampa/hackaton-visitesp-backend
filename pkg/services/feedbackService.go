package services

import (
	"github.com/turistaremsampa/hackaton-visitesp-backend/pkg/models"
	"github.com/turistaremsampa/hackaton-visitesp-backend/pkg/repositories"
)

func CreateFeedback(feedback *models.Feedback) error {
	return repositories.CreateFeedback(feedback)
}

func GetAllFeedbacks() ([]models.Feedback, error) {
	return repositories.GetAllFeedbacks()
}

func GetFeedbackByID(id string) (models.Feedback, error) {
	return repositories.GetFeedbackByID(id)
}
