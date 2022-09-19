package main

import (
	"fmt"
	"go-api/controllers/userController"
	userService "go-api/services/user"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", getRoot)

	r.GET("/users", userController.GetUsers)

	r.Run(":8888")

	users := userService.GetUsers()
	fmt.Println(users)
}

func getRoot(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World!",
	})
}
