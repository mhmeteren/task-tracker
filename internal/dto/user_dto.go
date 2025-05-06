package dto

import "task-tracker/internal/model"

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required,min=2"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func (r *CreateUserRequest) ToModel() model.User {
	return model.User{
		Name:  r.Name,
		Email: r.Email,
	}
}
