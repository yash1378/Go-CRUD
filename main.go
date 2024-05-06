package main

import (
	"myapp/controllers"
	"myapp/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	models.ConnectDatabase() // new!
	router.POST("/posts", controllers.CreatePost)
	router.GET("/posts/:id", controllers.FindPosts)
	router.PATCH("/posts/:id", controllers.UpdatePost)
	router.DELETE("/posts/:id", controllers.DeletePost)
	// ...

	router.Run("localhost:8080")
}
