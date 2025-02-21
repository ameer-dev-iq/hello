package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	ID    uint   `json:"id" gorm:"primary_key"`
	Title string `json:"title" validate:"required"`
}
