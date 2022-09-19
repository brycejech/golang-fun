package userController

import (
	userService "go-api/services/user"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users := userService.GetUsers()

	c.JSON(200, users)
}
