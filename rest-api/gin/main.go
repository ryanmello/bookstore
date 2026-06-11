package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ryanmello/bookstore/rest-api/gin/controllers"
	"github.com/ryanmello/bookstore/rest-api/gin/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main(){
	r := gin.Default()

	r.POST("/post", controllers.CreatePost)
	r.GET("/post", controllers.GetPosts)
	r.GET("/post/:id", controllers.GetPost)
	r.PUT("/post/:id", controllers.UpdatePost)
	r.DELETE("/post/:id", controllers.DeletePost)

	r.Run()
}
