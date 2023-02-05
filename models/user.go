package model

type User struct {
	Key   string `json:"key"`
	Name  string `json:"name" validate:"required"`
	Email *string `json:"email" validate:"required,email"`
	Permission string `json:"permission" validate:"required"`
}
