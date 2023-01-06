package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mohsenMj/go-starter-kit/app/initializers/database"
	"github.com/mohsenMj/go-starter-kit/app/models"
	"github.com/mohsenMj/go-starter-kit/app/repositories"
	"github.com/mohsenMj/go-starter-kit/app/requests"
	"github.com/mohsenMj/go-starter-kit/app/services"
)

type PostController interface {
	Index(ctx *gin.Context)
	Show(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type controller struct {
	service services.PostService
}

func NewPostController() PostController {
	service := services.NewPostService(repositories.NewPostRepository(database.DB))
	return &controller{
		service: service,
	}
}
func (c *controller) Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, c.service.Responses(c.service.All()))
}

func (c *controller) Show(ctx *gin.Context) {
	post := c.service.Get(ctx.Param("id"))
	if post.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Object Not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, c.service.Response(post))
}

func (c *controller) Create(ctx *gin.Context) {
	var reqeust requests.PostCreateRequest
	if err := ctx.ShouldBindJSON(&reqeust); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post := models.Post{Title: reqeust.Title, Body: reqeust.Body}
	c.service.Create(&post)
	ctx.JSON(http.StatusOK, c.service.Response(post))
}

func (c *controller) Update(ctx *gin.Context) {
	var input requests.PostUpdateRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post := c.service.Get(ctx.Param("id"))
	post.Title = input.Title
	post.Body = input.Body
	c.service.Update(&post)
	ctx.JSON(http.StatusOK, c.service.Response(post))
}

func (c *controller) Delete(ctx *gin.Context) {
	post := c.service.Get(ctx.Param("id"))
	if post.ID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Model not found"})
		return
	}
	c.service.Delete(ctx.Param("id"))
	ctx.Status(http.StatusOK)
}
