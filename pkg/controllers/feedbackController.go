package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/turistaremsampa/hackaton-visitesp-backend/pkg/models"
	"github.com/turistaremsampa/hackaton-visitesp-backend/pkg/services"
)

// CreateFeedback handles the creation of a new feedback
func CreateFeedback(c *gin.Context) {
	var feedback models.Feedback
	if err := c.ShouldBindJSON(&feedback); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreateFeedback(&feedback); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create feedback"})
		return
	}

	c.JSON(http.StatusOK, feedback)
}

// GetAllFeedbacks retrieves all feedbacks
func GetAllFeedbacks(c *gin.Context) {
	feedbacks, err := services.GetAllFeedbacks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve feedbacks"})
		return
	}
	c.JSON(http.StatusOK, feedbacks)
}

// GetFeedbackByID retrieves a feedback by its ID
func GetFeedbackByID(c *gin.Context) {
	id := c.Param("id")
	feedback, err := services.GetFeedbackByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Feedback not found"})
		return
	}
	c.JSON(http.StatusOK, feedback)
}
