package services

import (
	"github.com/mohsenMj/go-starter-kit/app/models"
	"github.com/mohsenMj/go-starter-kit/app/repositories"
	"github.com/mohsenMj/go-starter-kit/app/responses"
)

type PostService interface {
	Create(post *models.Post)
	Update(post *models.Post)
	Delete(id string)
	All() []models.Post
	Get(id string) models.Post
	Response(post models.Post) responses.PostResponse
}

type postService struct {
	rep repositories.PostRepository
}

func NewPostService(rep repositories.PostRepository) PostService {
	return &postService{
		rep: rep,
	}
}

func (s *postService) Create(post *models.Post) {
	s.rep.Create(post)
}

func (s *postService) Update(post *models.Post) {
	s.rep.Update(post)
}

func (s *postService) All() []models.Post {
	posts := s.rep.All()

	// var response []models.Post
	// for _, post := range posts {
	// 	response = append(response, models.Post{
	// 		ID:    post.ID,
	// 		Title: post.Title,
	// 		Body:  post.Body,
	// 	})
	// }
	return posts
}

func (s *postService) Get(id string) models.Post {
	var post models.Post
	post = s.rep.Get(id)
	return models.Post{
		ID:    post.ID,
		Title: post.Title,
		Body:  post.Body,
	}
}

func (s *postService) Delete(id string) {
	s.rep.Delete(id)
}

func (s *postService) Response(post models.Post) responses.PostResponse {
	return responses.PostResponse{
		ID:    post.ID,
		Title: post.Title,
		Body:  post.Body,
	}
}
