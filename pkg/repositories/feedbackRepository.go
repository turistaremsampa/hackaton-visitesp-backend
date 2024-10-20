package repositories

import (
	"github.com/turistaremsampa/hackaton-visitesp-backend/pkg/database"
	"github.com/turistaremsampa/hackaton-visitesp-backend/pkg/models"
)

func CreateFeedback(feedback *models.Feedback) error {
	return database.DB.Create(feedback).Error
}

func GetAllFeedbacks() ([]models.Feedback, error) {
	var feedbacks []models.Feedback
	if err := database.DB.Find(&feedbacks).Error; err != nil {
		return nil, err
	}
	return feedbacks, nil
}

func GetFeedbackByID(id string) (models.Feedback, error) {
	var feedback models.Feedback
	if err := database.DB.First(&feedback, id).Error; err != nil {
		return feedback, err
	}
	return feedback, nil
}
