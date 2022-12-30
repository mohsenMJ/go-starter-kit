package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	ID uint `gorm:"primarykey"`

	Title string `json:"title"`
	Body  string `json:"body"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
