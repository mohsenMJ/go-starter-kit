package models

import (
	"errors"
	"go-curd/app"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model

	Title string
	Body  string
}

func (p *Post) Create() (tx *gorm.DB) {
	return app.DB.Debug().Create(&p)
}

func (p *Post) Save() {
	app.DB.Debug().Save(&p)
}

func (p *Post) All() []Post {
	var posts []Post
	app.DB.Find(&posts)
	return posts
}

func (p *Post) Get(id string) error {
	app.DB.Find(&p, id)
	if p.ID == 0 {
		return errors.New("Object Not Found")
	}
	return nil
}

func (p *Post) GetOrFail(id string) error {
	app.DB.Find(&p, id)
	if p.ID == 0 {
		return errors.New("Object Not Found")
	}
	return nil
}

func (p *Post) Delete(id string) {
	app.DB.Delete(&p, id)
}
