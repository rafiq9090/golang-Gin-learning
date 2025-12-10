package handler

import (
	"go_project_Gin/internal/dto"
	"go_project_Gin/internal/service"
	"go_project_Gin/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllTasks(c *gin.Context) {
	task, err := service.Task.GetAllTasks(c.GetUint("user_id"))
	if err != nil {
		utils.JSONError(c, "Failed to get tasks", http.StatusInternalServerError, nil)
		return
	}
	utils.JSONSuccess(c, task)
}

func CreateTask(c *gin.Context) {
	var input dto.CreateTaskRequest
	userId := c.GetUint("user_id")

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.JSONError(c, "Invalid input", http.StatusBadRequest, nil)
		return
	}
	task, err := service.Task.CreateTask(userId, input.Title, input.Done)
	if err != nil {
		utils.JSONError(c, "Failed to create task", http.StatusInternalServerError, nil)
		return
	}
	utils.JSONSuccess(c, task)
}
