package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/turistaremsampa/hackaton-visitesp-backend/pkg/controllers"
)

func InitializeRoutes(r *gin.Engine) {
	r.POST("/users", controllers.CreateUser)
	r.GET("/users", controllers.GetAllUsers)
}
