package service

import (
	"context"
	"errors"
	"go_project_Gin/internal/model"
	"go_project_Gin/internal/repository"
)

type PostService struct{}

var Post = PostService{}

func (PostService) CreatePost(ctx context.Context, userId uint, caption, imageUrl string) (*model.Post, error) {
	if caption == "" && imageUrl == "" {
		return nil, errors.New("caption and imageUrl are required")
	}
	post := &model.Post{
		UserID:   userId,
		Caption:  caption,
		ImageUrl: imageUrl,
	}
	if err := repository.Post.Create(ctx, post); err != nil {
		return nil, err
	}
	return post, nil
}

func (PostService) GetPostsByUserId(ctx context.Context, userId uint) ([]model.Post, error) {
	return repository.Post.GetByUserId(ctx, userId)
}

func (PostService) GetAllPost(ctx context.Context) ([]model.Post, error) {
	return repository.Post.GetAllPost(ctx)
}
