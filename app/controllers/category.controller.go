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

type CategoryController interface {
	CategoryIndex(ctx *gin.Context)
	CategoryShow(ctx *gin.Context)
	CategoryCreate(ctx *gin.Context)
	CategoryUpdate(ctx *gin.Context)
	CategoryDelete(ctx *gin.Context)
}

type categoryController struct {
	service services.CategoryService
}

func NewCategoryController() CategoryController {
	service := services.NewCategoryService(repositories.NewCategoryRepository(database.DB))
	return &categoryController{
		service: service,
	}
}
func (c *categoryController) CategoryIndex(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, c.service.CategoryResponses(c.service.CategoryAll()))
}

func (c *categoryController) CategoryShow(ctx *gin.Context) {
	category := c.service.CategoryGet(ctx.Param("id"))
	if category.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Object Not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, c.service.CategoryResponse(category))
}

func (c *categoryController) CategoryCreate(ctx *gin.Context) {
	var reqeust requests.CategoryCreateRequest
	if err := ctx.ShouldBindJSON(&reqeust); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	category := models.Category{Title: reqeust.Title}
	c.service.CategoryCreate(&category)
	ctx.JSON(http.StatusOK, c.service.CategoryResponse(category))
}

func (c *categoryController) CategoryUpdate(ctx *gin.Context) {
	var input requests.CategoryUpdateRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	category := c.service.CategoryGet(ctx.Param("id"))
	category.Title = input.Title
	c.service.CategoryUpdate(&category)
	ctx.JSON(http.StatusOK, c.service.CategoryResponse(category))
}

func (c *categoryController) CategoryDelete(ctx *gin.Context) {
	category := c.service.CategoryGet(ctx.Param("id"))
	if category.ID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Model not found"})
		return
	}
	c.service.CategoryDelete(ctx.Param("id"))
	ctx.Status(http.StatusOK)
}
