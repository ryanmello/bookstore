package main

import (
	"github.com/ryanmello/bookstore/rest-api/gin/initializers"
	"github.com/ryanmello/bookstore/rest-api/gin/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
