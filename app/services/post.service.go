package services

import (
	"github.com/mohsenMj/go-starter-kit/app/models"
	"github.com/mohsenMj/go-starter-kit/app/repositories"
	"github.com/mohsenMj/go-starter-kit/app/responses"
)

type PostService interface {
	PostCreate(post *models.Post)
	PostUpdate(post *models.Post)
	PostDelete(id string)
	PostAll() []models.Post
	PostGet(id string) models.Post
	PostResponse(post models.Post) responses.PostResponse
	PostResponses(posts []models.Post) []responses.PostResponse
}

type postService struct {
	rep repositories.PostRepository
}

func NewPostService(rep repositories.PostRepository) PostService {
	return &postService{
		rep: rep,
	}
}

func (s *postService) PostCreate(post *models.Post) {
	s.rep.PostCreate(post)
}

func (s *postService) PostUpdate(post *models.Post) {
	s.rep.PostUpdate(post)
}

func (s *postService) PostAll() []models.Post {
	posts := s.rep.PostAll()
	return posts
}

func (s *postService) PostGet(id string) models.Post {
	var post models.Post
	post = s.rep.PostGet(id)
	return models.Post{
		ID:    post.ID,
		Title: post.Title,
		Body:  post.Body,
	}
}

func (s *postService) PostDelete(id string) {
	s.rep.PostDelete(id)
}

func (s *postService) PostResponse(post models.Post) responses.PostResponse {
	return responses.PostResponse{
		ID:    post.ID,
		Title: post.Title,
		Body:  post.Body,
	}
}

func (s *postService) PostResponses(posts []models.Post) []responses.PostResponse {
	var response []responses.PostResponse
	for _, post := range posts {
		response = append(response, s.PostResponse(post))
	}
	return response
}
