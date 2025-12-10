package service

import (
	"go_project_Gin/internal/model"
	"go_project_Gin/internal/repository"
)

type TaskService struct{}

var Task = TaskService{}

func (TaskService) GetAllTasks(userId uint) ([]model.Task, error) {
	return repository.Task.GetByUserID(userId)
}

func (TaskService) CreateTask(userId uint, title string, done bool) (model.Task, error) {
	task := model.Task{
		UserID: userId,
		Title:  title,
		Done:   done,
	}
	err := repository.Task.CreateTask(task)
	return task, err
}
