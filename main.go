package main

import (
	"github.com/turistaremsampa/hackaton-visitesp-backend/pkg/routes"

	"github.com/gin-gonic/gin"
	"github.com/turistaremsampa/hackaton-visitesp-backend/pkg/database"
)

func main() {
	// Initialize Gin framework
	r := gin.Default()

	// Connect to the database
	database.ConnectDB()

	// Set up routes
	routes.InitializeRoutes(r)

	// Start the server
	r.Run(":8080")
}
