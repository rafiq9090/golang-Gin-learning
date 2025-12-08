package handler

import (
	"go_project_Gin/internal/database"
	"go_project_Gin/internal/dto"
	"go_project_Gin/internal/model"
	"go_project_Gin/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	var input dto.CreateTaskRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.JSONError(c, "Invalid input", http.StatusBadRequest, nil)
		return
	}
	task := model.Task{
		Title: input.Title,
		Done:  input.Done,
	}
	database.DB.Create(&task)
	utils.JSONSuccess(c, task)
}
