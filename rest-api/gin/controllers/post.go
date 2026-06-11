package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ryanmello/bookstore/rest-api/gin/initializers"
	"github.com/ryanmello/bookstore/rest-api/gin/models"
)

func CreatePost(c *gin.Context) {
	// get data from request body
	var body struct {
		Title string
		Body string
	}

	c.Bind(&body)

	// create post object
	post := models.Post{Title: body.Title, Body: body.Body}

	// insert into database
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func GetPosts(c *gin.Context){
	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func GetPost(c *gin.Context){
	// get id from url
	id := c.Param("id")

	var post models.Post
	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"post": post,
	})
}

func UpdatePost(c *gin.Context){
	id := c.Param("id")

	var body struct {
		Title string
		Body string
	}

	c.Bind(&body)

	// create post object
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		c.Status(400)
		return
	}

	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Title})

	c.JSON(200, gin.H{
		"post": post,
	})
}

func DeletePost(c *gin.Context){
	id := c.Param("id")
	initializers.DB.Delete(&models.Post{}, id)
	c.Status(200)
}
