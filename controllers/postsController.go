package controllers

import (
	"go-curd/app"
	"go-curd/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func PostsIndex(c *gin.Context) {
	var posts []models.Post
	app.DB.Find(&posts)
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	app.DB.Find(&post, id)
	c.JSON(200, gin.H{
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
	result := app.DB.Create(&post)
	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})

}

func PostUpdate(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	app.DB.Find(&post, id)

	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)

	//First Approach
	app.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	//Second Approach
	// post.Title = body.Title
	// post.Body = body.Title
	// app.DB.Save(&post)

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	id := c.Param("id")
	app.DB.Delete(&models.Post{}, id)

	c.Status(http.StatusOK)
}
