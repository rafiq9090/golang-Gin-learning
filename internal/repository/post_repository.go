package repository

import (
	"context"
	"go_project_Gin/internal/database"
	"go_project_Gin/internal/model"
)

type PostRepository struct{}

var Post PostRepository

func (PostRepository) Create(ctx context.Context, post *model.Post) error {
	return database.DB.WithContext(ctx).Create(post).Error
}

func (PostRepository) GetByUserId(ctx context.Context, userId uint) ([]model.Post, error) {
	var posts []model.Post
	err := database.DB.WithContext(ctx).Where("user_id = ?", userId).
		Order("created_at desc").Find(&posts).Error
	return posts, err
}

func (PostRepository) GetAllPost(ctx context.Context) ([]model.Post, error) {
	var posts []model.Post
	err := database.DB.WithContext(ctx).Order("created_at desc").Find(&posts).Error
	return posts, err
}
