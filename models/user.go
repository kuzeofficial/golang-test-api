package model

type User struct {
	Key   string `json:"key"`
	Name  string `json:"name" validate:"required"`
	Title string `json:"title" validate:"required"`
}
