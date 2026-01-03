package service

import (
	"context"
	"errors"
	"go_project_Gin/internal/model"
	"go_project_Gin/internal/repository"
)

type TaskService struct{}

var Task = TaskService{}

func (TaskService) GetAllTasks(ctx context.Context) ([]model.Task, error) {
	return repository.Task.GetAllTasks(ctx)
}

func (TaskService) CreateTask(ctx context.Context, userId uint, title string, done bool) (model.Task, error) {
	task := model.Task{
		UserID: userId,
		Title:  title,
		Done:   done,
	}
	err := repository.Task.CreateTask(ctx, &task)
	return task, err
}

func (TaskService) UpdateTask(ctx context.Context, id uint64, userId uint, title string, done bool) (model.Task, error) {
	existingTask, err := repository.Task.GetByID(ctx, uint(id), userId)
	if err != nil {
		return model.Task{}, err
	}

	existingTask.Title = title
	existingTask.Done = done

	err = repository.Task.UpdateTask(ctx, existingTask)
	if err != nil {
		return model.Task{}, err
	}

	return *existingTask, nil
}

// func (TaskService) UpdateTaskModel(ctx context.Context, task *model.Task) error {
// 	task, err := repository.Task.GetByID(ctx, task.ID, task.UserID)
// 	if err != nil {
// 		return errors.New("task not found or unauthorized")
// 	}
// 	task.Title = task.Title
// 	task.Done = task.Done
// 	return repository.Task.UpdateTask(ctx, task)

// }

func (TaskService) DeleteTask(ctx context.Context, id uint, userId uint) error {
	_, err := repository.Task.GetByID(ctx, id, userId)
	if err != nil {
		return errors.New("task not found or unauthorized")
	}
	return repository.Task.DeleteTask(ctx, id)

}
