package repositories

import (
	"github.com/mohsenMj/go-starter-kit/app/models"
	"gorm.io/gorm"
)

type PostRepository interface {
	Create(post *models.Post) *models.Post
	Update(post *models.Post) *models.Post
	Delete(id string)
	All() []models.Post
	Get(id string) models.Post
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{
		db: db,
	}
}

func (r *postRepository) Create(post *models.Post) *models.Post {
	r.db.Debug().Create(&post)
	return post
}

func (r *postRepository) Update(post *models.Post) *models.Post {
	r.db.Save(&post)
	return post
}

func (r *postRepository) Delete(id string) {
	r.db.Debug().Delete(&models.Post{}, id)
}

func (r *postRepository) All() []models.Post {
	var posts []models.Post
	// database.DB.Debug().Find(&posts)
	r.db.Debug().Find(&posts)
	return posts
}

func (r *postRepository) Get(id string) models.Post {
	var post models.Post
	r.db.Debug().Find(&post, id)
	return post
}
