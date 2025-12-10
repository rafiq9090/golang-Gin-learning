package repository

import (
	"errors"
	"go_project_Gin/internal/database"
	"go_project_Gin/internal/model"
)

type AuthRepository struct{}

var Auth = AuthRepository{}

func (AuthRepository) Register(user *model.User) error {
	return database.DB.Create(user).Error
}

func (AuthRepository) Login(email string, password string) (*model.User, error) {
	var user model.User
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	if user.Password != password {
		return nil, errors.New("invalid password")
	}
	return &user, nil
}

func (AuthRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
