package models

import (
	"go-curd/app"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title string
	Body  string
}

// find a way to simplify this
func (p *Post) Save() {
	app.DB.Save(&p)
}
