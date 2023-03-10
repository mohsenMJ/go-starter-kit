package models

import (
	"time"
)

type Post struct {
	ID uint `gorm:"primarykey"`

	Title string `json:"title"`
	Body  string `json:"body"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
