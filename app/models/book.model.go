package models

import "time"

type Book struct {
	ID          uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Title       string `gorm:"type:varchar(255)" json:"title"`
	Description string `gorm:"type:text" json:"description"`
	UserID      uint64 `gorm:"not null" json:"-"`
	User        User   `gorm:"foreignkey:UserID;constaint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
