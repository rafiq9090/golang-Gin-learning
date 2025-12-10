package service

import (
	"errors"
	"go_project_Gin/internal/model"
	"go_project_Gin/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

var JWT_SECRET = "secret"

type AuthService struct{}

var Auth = AuthService{}

func (AuthService) Register(user *model.User) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashed)
	return repository.Auth.Register(user)
}

func (AuthService) Login(email string, password string) (*model.User, error) {
	user, err := repository.Auth.FindByEmail(email)
	if err != nil {
		return nil, errors.New("Invalid credentials")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("Invalid credentials")
	}
	return user, nil
}
