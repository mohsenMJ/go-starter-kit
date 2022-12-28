package controllers

import (
	"go-curd/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func PostsIndex(c *gin.Context) {
	var post models.Post
	c.JSON(200, gin.H{
		"posts": post.All(),
	})
}

func PostsShow(c *gin.Context) {
	post := models.Post{}
	err := post.GetOrFail(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": strings.Split(err.Error(), ","),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})

}

func PostCreate(c *gin.Context) {
	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}
	result := post.Create()
	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})

}

func PostUpdate(c *gin.Context) {
	var post models.Post
	post.Get(c.Param("id"))

	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)

	post.Title = body.Title
	post.Body = body.Body
	post.Save()

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	post := models.Post{}
	post.Delete(c.Param("id"))

	c.Status(http.StatusOK)
}
