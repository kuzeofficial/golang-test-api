package model

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name" validate:"required"`
	Title string `json:"title" validate:"required"`
}