package handler

import (
	"go_project_Gin/internal/config"
	"go_project_Gin/internal/dto"
	"go_project_Gin/internal/model"
	"go_project_Gin/internal/service"
	"go_project_Gin/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

var JWT_SECRET = []byte(config.App.JWTSecret)

func RegisterHandler(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.JSONError(c, "Invalid request", http.StatusBadRequest, nil)
		return
	}
	if errs := dto.ValidateRegister(req); errs != nil {
		utils.JSONError(c, "Validation failed", http.StatusBadRequest, errs)
		return
	}

	user := model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := service.Auth.Register(&user); err != nil {
		utils.JSONError(c, "Email already exists", http.StatusBadRequest, nil)
		return
	}

	user.Password = ""

	utils.JSONSuccess(c, map[string]any{
		"message": "User registered successfully",
		"user":    user,
	})
}

func LoginHandler(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.JSONError(c, "Invalid request", http.StatusBadRequest, nil)
		return
	}
	if errs := dto.ValidateLogin(req); errs != nil {
		utils.JSONError(c, "Validation failed", http.StatusBadRequest, errs)
		return
	}

	user, token, err := service.Auth.Login(req.Email, req.Password)
	if err != nil {
		utils.JSONError(c, "Invalid email or password", http.StatusUnauthorized, nil) // ← 401
		return
	}

	user.Password = ""

	utils.JSONSuccess(c, map[string]any{
		"message": "Login successful",
		"token":   token,
		"user": map[string]any{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}
