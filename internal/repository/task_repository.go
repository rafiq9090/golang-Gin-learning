package repository

import (
	"context"
	"go_project_Gin/internal/database"
	"go_project_Gin/internal/model"
)

type TaskRepository struct{}

var Task = TaskRepository{}

func (TaskRepository) GetAllTasks(ctx context.Context) ([]model.Task, error) {
	var task []model.Task

	// result := database.DB.Find(&task)
	// if result.Error != nil {
	// 	return nil, result.Error
	// }
	// return task, nil

	result := database.DB.WithContext(ctx).Find(&task)
	return task, result.Error

}
func (TaskRepository) GetByUserID(ctx context.Context, userId uint) ([]model.Task, error) {
	var task []model.Task
	result := database.DB.WithContext(ctx).Where("user_id = ?", userId).Find(&task)
	// result := database.DB.Where("user_id = ?", userId).Find(&task)
	return task, result.Error
}

func (TaskRepository) CreateTask(ctx context.Context, task *model.Task) error {
	// result := database.DB.Create(&task)
	// if result.Error != nil {
	// 	return task, result.Error
	// }
	// return task, nil
	return database.DB.WithContext(ctx).Create(task).Error
}

func (TaskRepository) GetByID(ctx context.Context, id uint, userId uint) (*model.Task, error) {
	var task model.Task
	err := database.DB.WithContext(ctx).Where("id = ? AND user_id = ?", id, userId).First(&task).Error

	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (TaskRepository) UpdateTask(ctx context.Context, task *model.Task) error {
	return database.DB.WithContext(ctx).Save(task).Error
}

func (TaskRepository) DeleteTask(ctx context.Context, id uint) error {
	return database.DB.WithContext(ctx).Delete(&model.Task{}, id).Error
}
