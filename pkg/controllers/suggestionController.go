package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/turistaremsampa/hackaton-visitesp-backend/pkg/models"
	"github.com/turistaremsampa/hackaton-visitesp-backend/pkg/services"
)

// CreateSuggestion handles the creation of a new suggestion
func CreateSuggestion(c *gin.Context) {
	var suggestion models.Suggestion
	if err := c.ShouldBindJSON(&suggestion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreateSuggestion(&suggestion); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create suggestion"})
		return
	}

	c.JSON(http.StatusOK, suggestion)
}

// GetAllSuggestions retrieves all suggestions
func GetAllSuggestions(c *gin.Context) {
	suggestions, err := services.GetAllSuggestions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve suggestions"})
		return
	}
	c.JSON(http.StatusOK, suggestions)
}

// GetSuggestionByID retrieves a suggestion by its ID
func GetSuggestionByID(c *gin.Context) {
	id := c.Param("id")
	suggestion, err := services.GetSuggestionByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Suggestion not found"})
		return
	}
	c.JSON(http.StatusOK, suggestion)
}
