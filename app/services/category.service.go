package services

import (
	"github.com/mohsenMj/go-starter-kit/app/models"
	"github.com/mohsenMj/go-starter-kit/app/repositories"
	"github.com/mohsenMj/go-starter-kit/app/responses"
)

type CategoryService interface {
	CategoryCreate(category *models.Category)
	CategoryUpdate(category *models.Category)
	CategoryDelete(id string)
	CategoryAll() []models.Category
	CategoryGet(id string) models.Category
	CategoryResponse(category models.Category) responses.CategoryResponse
	CategoryResponses(categorys []models.Category) []responses.CategoryResponse
}

type categoryService struct {
	rep repositories.CategoryRepository
}

func NewCategoryService(rep repositories.CategoryRepository) CategoryService {
	return &categoryService{
		rep: rep,
	}
}

func (s *categoryService) CategoryCreate(category *models.Category) {
	s.rep.CategoryCreate(category)
}

func (s *categoryService) CategoryUpdate(category *models.Category) {
	s.rep.CategoryUpdate(category)
}

func (s *categoryService) CategoryAll() []models.Category {
	categorys := s.rep.CategoryAll()
	return categorys
}

func (s *categoryService) CategoryGet(id string) models.Category {
	var category models.Category
	category = s.rep.CategoryGet(id)
	return models.Category{
		ID:    category.ID,
		Title: category.Title,
	}
}

func (s *categoryService) CategoryDelete(id string) {
	s.rep.CategoryDelete(id)
}

func (s *categoryService) CategoryResponse(category models.Category) responses.CategoryResponse {
	return responses.CategoryResponse{
		ID:    category.ID,
		Title: category.Title,
	}
}

func (s *categoryService) CategoryResponses(categorys []models.Category) []responses.CategoryResponse {
	var response []responses.CategoryResponse
	for _, category := range categorys {
		response = append(response, s.CategoryResponse(category))
	}
	return response
}
