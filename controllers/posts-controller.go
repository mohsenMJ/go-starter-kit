package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mohsenMj/go-starter-kit/models"
	"github.com/mohsenMj/go-starter-kit/services"
)

type CreatePostInput struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
}

type PostController interface {
	Index(ctx *gin.Context) []models.Post
	Show(ctx *gin.Context) models.Post
	Create(ctx *gin.Context) models.Post
	Update(ctx *gin.Context) models.Post
	Delete(ctx *gin.Context) models.Post
}

type controller struct {
	service services.PostService
}

func New(service services.PostService) controller {
	return controller{
		service: service,
	}
}

func (c *controller) Index(ctx *gin.Context) []models.Post {
	return c.service.All()
}

func (c *controller) Show(ctx *gin.Context, id string) models.Post {
	return c.service.Get(ctx.Param("id"))
}

func (c *controller) Create(ctx *gin.Context) models.Post {
	var post models.Post
	var input CreatePostInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return post
	}
	post = models.Post{Title: input.Title, Body: input.Body}
	c.service.Create(&post)
	return post
}

func (c *controller) Update(ctx *gin.Context) models.Post {
	post := c.service.Get(ctx.Param("id"))
	ctx.BindJSON(&post)
	c.service.Save(post)
	return post
}

func (c *controller) Delete(ctx *gin.Context) {
	c.service.Delete(ctx.Param("id"))
	ctx.Status(http.StatusOK)
}
