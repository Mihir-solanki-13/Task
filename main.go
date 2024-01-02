package main

import (
	"github.com/gin-gonic/gin"
	"TASK/internal/handlers"
	"TASK/internal/database"

)

func main() {
	// Initialize the database
	database.InitDatabase()

	// Create a new Gin router
	router := gin.Default()

	// Define API routes
	router.GET("/tasks", handlers.GetAllTasks)
	router.GET("/tasks/:id", handlers.GetTaskByID)
	router.POST("/tasks", handlers.CreateTask)
	router.PUT("/tasks/:id", handlers.UpdateTask)
	router.DELETE("/tasks/:id", handlers.DeleteTask)

	// Run the server
	router.Run(":8080")
}
