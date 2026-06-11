package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryanmello/bookstore/rest-api/gin/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main(){
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H {
			"message": "pong",
		})
	})

	r.Run()
}
