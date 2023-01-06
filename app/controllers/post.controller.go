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
	PostIndex(ctx *gin.Context)
	PostShow(ctx *gin.Context)
	PostCreate(ctx *gin.Context)
	PostUpdate(ctx *gin.Context)
	PostDelete(ctx *gin.Context)
}

type postController struct {
	service services.PostService
}

func NewPostController() PostController {
	service := services.NewPostService(repositories.NewPostRepository(database.DB))
	return &postController{
		service: service,
	}
}
func (c *postController) PostIndex(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, c.service.PostResponses(c.service.PostAll()))
}

func (c *postController) PostShow(ctx *gin.Context) {
	post := c.service.PostGet(ctx.Param("id"))
	if post.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Object Not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, c.service.PostResponse(post))
}

func (c *postController) PostCreate(ctx *gin.Context) {
	var reqeust requests.PostCreateRequest
	if err := ctx.ShouldBindJSON(&reqeust); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post := models.Post{Title: reqeust.Title, Body: reqeust.Body}
	c.service.PostCreate(&post)
	ctx.JSON(http.StatusOK, c.service.PostResponse(post))
}

func (c *postController) PostUpdate(ctx *gin.Context) {
	var input requests.PostUpdateRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post := c.service.PostGet(ctx.Param("id"))
	post.Title = input.Title
	post.Body = input.Body
	c.service.PostUpdate(&post)
	ctx.JSON(http.StatusOK, c.service.PostResponse(post))
}

func (c *postController) PostDelete(ctx *gin.Context) {
	post := c.service.PostGet(ctx.Param("id"))
	if post.ID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Model not found"})
		return
	}
	c.service.PostDelete(ctx.Param("id"))
	ctx.Status(http.StatusOK)
}
