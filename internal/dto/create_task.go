package dto

import "github.com/go-playground/validator/v10"

type CreateTaskRequest struct {
	Title string `json:"title" validate:"required min=3 max=100"`
	Done  bool   `json:"done"`
}

func ValidateCreateTask(req CreateTaskRequest) map[string]string {
	err := validator.New().Struct(req)
	if err == nil {
		return nil
	}
	errors := make(map[string]string)
	for _, err := range err.(validator.ValidationErrors) {
		field := err.Field()
		tag := err.Tag()

		if tag == "required" {
			errors[field] = field + " is required"
		} else if tag == "min" {
			errors[field] = field + " must be at least 2 characters"
		} else if tag == "max" {
			errors[field] = field + "must be at most 100 characters"
		}
	}
	return errors
}
