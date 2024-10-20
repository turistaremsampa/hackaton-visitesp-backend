package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/turistaremsampa/hackaton-visitesp-backend/pkg/controllers"
)

func InitializeRoutes(r *gin.Engine) {

	r.POST("/users", controllers.CreateUser)
	r.GET("/users", controllers.GetAllUsers)

	// Locations feedbacks
	r.POST("/feedbacks", controllers.CreateFeedback)
	r.GET("/feedbacks", controllers.GetAllFeedbacks)
	r.GET("/feedbacks/:id", controllers.GetFeedbackByID)

	// Locations suggestions
	r.POST("/suggestions", controllers.CreateSuggestion)
	r.GET("/suggestions", controllers.GetAllSuggestions)
	r.GET("/suggestions/:id", controllers.GetSuggestionByID)
}
