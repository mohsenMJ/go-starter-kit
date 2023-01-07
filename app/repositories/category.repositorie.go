package repositories

import (
	"github.com/mohsenMj/go-starter-kit/app/models"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	CategoryCreate(category *models.Category) *models.Category
	CategoryUpdate(category *models.Category) *models.Category
	CategoryDelete(id string)
	CategoryAll() []models.Category
	CategoryGet(id string) models.Category
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}

func (r *categoryRepository) CategoryCreate(category *models.Category) *models.Category {
	r.db.Debug().Create(&category)
	return category
}

func (r *categoryRepository) CategoryUpdate(category *models.Category) *models.Category {
	r.db.Save(&category)
	return category
}

func (r *categoryRepository) CategoryDelete(id string) {
	r.db.Debug().Delete(&models.Category{}, id)
}

func (r *categoryRepository) CategoryAll() []models.Category {
	var categorys []models.Category
	r.db.Debug().Find(&categorys)
	return categorys
}

func (r *categoryRepository) CategoryGet(id string) models.Category {
	var category models.Category
	r.db.Debug().Find(&category, id)
	return category
}
