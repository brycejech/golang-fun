package main

import (
	"go-api/controllers/userController"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", getRoot)

	r.GET("/users", userController.GetUsers)
	r.POST("/users", userController.CreateUser)

	r.Run(":8888")
}

func getRoot(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World!",
	})
}
