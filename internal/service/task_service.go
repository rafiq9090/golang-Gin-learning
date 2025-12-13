package service

import (
	"context"
	"go_project_Gin/internal/model"
	"go_project_Gin/internal/repository"
)

type TaskService struct{}

var Task = TaskService{}

func (TaskService) GetAllTasks(ctx context.Context, userId uint) ([]model.Task, error) {
	return repository.Task.GetByUserID(ctx, userId)
}

func (TaskService) CreateTask(ctx context.Context, userId uint, title string, done bool) (model.Task, error) {
	task := model.Task{
		UserID: userId,
		Title:  title,
		Done:   done,
	}
	err := repository.Task.CreateTask(ctx, task)
	return task, err
}
