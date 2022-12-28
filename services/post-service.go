package services

import (
	"log"

	"github.com/mohsenMj/go-starter-kit/app"
	"github.com/mohsenMj/go-starter-kit/models"
)

type PostService interface {
	Create(*models.Post)
	Save(*models.Post)
	All() []models.Post
	Get(id string) models.Post
	Delete(id string)
}

type postService struct{}

func NewPostService() PostService {
	return &postService{}
}

func (s *postService) Create(post *models.Post) {
	app.DB.Debug().Create(&post)
}

func (s *postService) Save(post *models.Post) {
	app.DB.Debug().Save(&post)
}

func (s *postService) All() []models.Post {
	var posts []models.Post
	app.DB.Debug().Find(&posts)
	log.Println("inside the All service")

	return posts
}

func (s *postService) Get(id string) models.Post {
	var post models.Post
	app.DB.Debug().Find(&post, id)
	return post
}

func (s *postService) Delete(id string) {
	app.DB.Debug().Delete(&models.Post{}, id)
}
