package handler

import (
	"go_project_Gin/internal/dto"
	"go_project_Gin/internal/notification"
	"go_project_Gin/internal/service"
	"go_project_Gin/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllTasks(c *gin.Context) {
	ctx := c.Request.Context()
	task, err := service.Task.GetAllTasks(ctx)
	if err != nil {
		utils.JSONError(c, "Failed to get tasks", http.StatusInternalServerError, nil)
		return
	}

	utils.JSONSuccess(c, task)
}

func CreateTask(c *gin.Context) {
	var input dto.CreateTaskRequest
	ctx := c.Request.Context()
	userId := c.GetUint("user_id")

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.JSONError(c, "Invalid input", http.StatusBadRequest, nil)
		return
	}
	task, err := service.Task.CreateTask(ctx, userId, input.Title, input.Done)
	if err != nil {
		utils.JSONError(c, "Failed to create task", http.StatusInternalServerError, nil)
		return
	}

	user, err := service.Auth.GetUserById(ctx, userId)
	if err != nil {
		utils.JSONError(c, "Failed to get user", http.StatusInternalServerError, nil)
		return
	}
	go notification.SendTaskNotification(user.Email, user.ID, task.ID, "created")

	utils.JSONSuccess(c, task)
}

func UpdateTask(c *gin.Context) {
	ctx := c.Request.Context()
	userId := c.GetUint("user_id")
	id := c.Param("id")

	taskId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		utils.JSONError(c, "Invalid task ID", http.StatusBadRequest, nil)
		return
	}
	var input dto.CreateTaskRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.JSONError(c, "Invalid input", http.StatusBadRequest, nil)
		return
	}
	task, err := service.Task.UpdateTask(ctx, taskId, userId, input.Title, input.Done)
	if err != nil {
		utils.JSONError(c, "Failed to update task", http.StatusInternalServerError, nil)
		return
	}

	user, err := service.Auth.GetUserById(ctx, userId)
	if err != nil {
		utils.JSONError(c, "Failed to get user", http.StatusInternalServerError, nil)
		return
	}
	go notification.SendTaskNotification(user.Email, user.ID, task.ID, "updated")

	utils.JSONSuccess(c, task)

}

func DeleteTask(c *gin.Context) {
	ctx := c.Request.Context()

	userID := c.GetUint("user_id")
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.JSONError(c, "Invalid task ID", http.StatusBadRequest, nil)
		return
	}

	if err := service.Task.DeleteTask(ctx, uint(id), userID); err != nil {
		utils.JSONError(c, err.Error(), http.StatusNotFound, nil)
		return
	}

	user, err := service.Auth.GetUserById(ctx, userID)
	if err != nil {
		utils.JSONError(c, "Failed to get user", http.StatusInternalServerError, nil)
		return
	}
	go notification.SendTaskNotification(user.Email, user.ID, uint(id), "deleted")

	c.Status(http.StatusNoContent)
}
