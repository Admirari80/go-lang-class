package main

import (
	"siashish/blogpost/controllers"
	"siashish/blogpost/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	models.ConnectDatabase()

	router.POST("/posts", controllers.CreatePost)

	router.Run("localhost:9000")
}
