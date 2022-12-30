package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mohsenMj/go-starter-kit/app/initializers/database"
	"github.com/mohsenMj/go-starter-kit/app/inputs"
	"github.com/mohsenMj/go-starter-kit/app/models"
	"github.com/mohsenMj/go-starter-kit/app/repositories"
	"github.com/mohsenMj/go-starter-kit/app/responses"
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
	posts := c.service.All()
	var response []responses.PostResponse
	for _, post := range posts {
		response = append(response, c.service.Response(post))
	}
	ctx.JSON(http.StatusOK, response)
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
	var input inputs.PostCreateInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post := models.Post{Title: input.Title, Body: input.Body}
	c.service.Create(&post)
	ctx.JSON(http.StatusOK, c.service.Response(post))
}

func (c *controller) Update(ctx *gin.Context) {
	var input inputs.PostUpdateInput
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
