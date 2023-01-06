package repositories

import (
	"github.com/mohsenMj/go-starter-kit/app/models"
	"gorm.io/gorm"
)

type PostRepository interface {
	PostCreate(post *models.Post) *models.Post
	PostUpdate(post *models.Post) *models.Post
	PostDelete(id string)
	PostAll() []models.Post
	PostGet(id string) models.Post
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{
		db: db,
	}
}

func (r *postRepository) PostCreate(post *models.Post) *models.Post {
	r.db.Debug().Create(&post)
	return post
}

func (r *postRepository) PostUpdate(post *models.Post) *models.Post {
	r.db.Save(&post)
	return post
}

func (r *postRepository) PostDelete(id string) {
	r.db.Debug().Delete(&models.Post{}, id)
}

func (r *postRepository) PostAll() []models.Post {
	var posts []models.Post
	r.db.Debug().Find(&posts)
	return posts
}

func (r *postRepository) PostGet(id string) models.Post {
	var post models.Post
	r.db.Debug().Find(&post, id)
	return post
}
