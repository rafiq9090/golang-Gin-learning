package service

import (
	"go_project_Gin/internal/model"
	"go_project_Gin/internal/repository"
)

type TaskService struct{}

var Task = TaskService{}

func (TaskService) GetAllTasks() ([]model.Task, error) {
	return repository.Task.GetAllTasks()
}

func (TaskService) CreateTask(title string, done bool) (model.Task, error) {
	task := model.Task{
		Title: title,
		Done:  done,
	}
	err := repository.Task.CreateTask(task)
	return task, err
}
