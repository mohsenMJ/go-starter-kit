package models

import (
	"time"
)

type Category struct {
	ID uint `gorm:"primarykey"`

	Title string `json:"title"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
