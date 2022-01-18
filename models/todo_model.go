package models

import "github.com/jinzhu/gorm"

type Todo struct {
	gorm.Model
	Title       string `json:"title" example:"Learning Golang"`
	Description string `json:"description" example:"Start from module 7"`
}

type TodoPayload struct {
	Title       string `json:"title" example:"Learning Golang"`
	Description string `json:"description" example:"Start from module 7"`
}
