package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mohsenMj/go-starter-kit/app/services"
	"github.com/mohsenMj/go-starter-kit/database/models"
)

type createPostInput struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
}

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
	service := services.NewPostService()
	return &controller{
		service: service,
	}
}
func (c *controller) Index(ctx *gin.Context) {
	log.Println("inside the index controller")
	ctx.JSON(http.StatusOK, c.service.All())
}

func (c *controller) Show(ctx *gin.Context) {
	post := c.service.Get(ctx.Param("id"))
	if post.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Object Not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, c.service.Get(ctx.Param("id")))
}

func (c *controller) Create(ctx *gin.Context) {
	var post models.Post
	var input createPostInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post = models.Post{Title: input.Title, Body: input.Body}
	c.service.Create(&post)
	ctx.JSON(http.StatusOK, post)
}

func (c *controller) Update(ctx *gin.Context) {
	post := c.service.Get(ctx.Param("id"))
	ctx.BindJSON(&post)
	c.service.Save(&post)
	ctx.JSON(http.StatusOK, post)
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
