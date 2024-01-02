package handlers

import (
	"github.com/gin-gonic/gin"
	"TASK/internal/database"
	"TASK/internal/models"
	"net/http"
)

// GetAllTasks retrieves all tasks from the database
func GetAllTasks(c *gin.Context) {
	var tasks []models.Task
	database.DB.Find(&tasks)
	c.JSON(http.StatusOK, tasks)
}

// GetTaskByID retrieves a task by its ID from the database
func GetTaskByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var task models.Task

	if err := database.DB.First(&task, id).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, task)
}

// CreateTask creates a new task and stores it in the database
func CreateTask(c *gin.Context) {
	var task models.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&task)
	c.JSON(http.StatusCreated, task)
}

// UpdateTask updates an existing task in the database
func UpdateTask(c *gin.Context) {
	id := c.Params.ByName("id")
	var task models.Task

	if err := database.DB.First(&task, id).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&task)
	c.JSON(http.StatusOK, task)
}

// DeleteTask deletes a task by its ID from the database
func DeleteTask(c *gin.Context) {
	id := c.Params.ByName("id")
	var task models.Task

	if err := database.DB.First(&task, id).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	database.DB.Delete(&task)
	c.Status(http.StatusNoContent)
}
