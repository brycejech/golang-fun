package main

import (
	"go-api/controllers"

	"github.com/gin-gonic/gin"
)

var (
	userController = controllers.NewUserController()
)

func main() {
	r := gin.Default()

	r.GET("/", getRoot)

	r.GET("/users", userController.GetUsers)
	r.POST("/users", userController.CreateUser)
	r.GET("/users/:id", userController.GetUser)

	r.Run(":8888")
}

func getRoot(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World!",
	})
}
