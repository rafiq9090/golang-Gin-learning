package service

import (
	"context"
	"errors"
	"go_project_Gin/internal/config"
	"go_project_Gin/internal/dto"
	"go_project_Gin/internal/model"
	"go_project_Gin/internal/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct{}

var Auth = AuthService{}

func (AuthService) Register(user *model.User) error {
	if _, err := repository.Auth.FindByEmail(user.Email); err == nil {
		return errors.New("Email already exists")
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashed)
	return repository.Auth.Register(user)
}

func (AuthService) Login(email string, password string) (*model.User, string, error) {
	user, err := repository.Auth.FindByEmail(email)
	if err != nil {
		return nil, "", errors.New("Invalid credentials")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, "", errors.New("Invalid credentials")
	}
	claims := dto.JWTClaim{
		UserId: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.App.JWTSecret))
	if err != nil {
		return nil, "", err
	}
	return user, tokenString, nil
}

func (AuthService) GetUserById(ctx context.Context, id uint) (*model.User, error) {
	return repository.Auth.FindByID(id)
}
