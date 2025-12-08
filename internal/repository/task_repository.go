package repository

import (
	"go_project_Gin/internal/database"
	"go_project_Gin/internal/model"
)

type TaskRepository struct{}

var Task = TaskRepository{}

func (TaskRepository) GetAllTasks() ([]model.Task, error) {
	var task []model.Task

	// result := database.DB.Find(&task)
	// if result.Error != nil {
	// 	return nil, result.Error
	// }
	// return task, nil

	result := database.DB.Find(&task)
	return task, result.Error

}

func (TaskRepository) CreateTask(task model.Task) error {
	// result := database.DB.Create(&task)
	// if result.Error != nil {
	// 	return task, result.Error
	// }
	// return task, nil
	return database.DB.Create(&task).Error
}
