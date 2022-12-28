package services

import (
	"github.com/mohsenMj/go-starter-kit/app"
	"github.com/mohsenMj/go-starter-kit/models"
)

type PostService interface {
	Create(*models.Post) models.Post
	Save(models.Post) models.Post
	All() []models.Post
	Get(id string) models.Post
	Delete(id string)
}

// this is only to be able to access the real methods
type service struct {
}

func (s *service) Create(post *models.Post) {
	app.DB.Debug().Create(&post)
}

func (s *service) Save(post models.Post) models.Post {
	app.DB.Debug().Save(&post)
	return post
}

func (s *service) All() []models.Post {
	var posts []models.Post
	app.DB.Find(&posts)
	return posts
}

func (s *service) Get(id string) models.Post {
	var post models.Post
	app.DB.Find(&post, id)
	return post
}

func (s *service) Delete(id string) {
	app.DB.Debug().Delete(&models.Post{}, id)
}
