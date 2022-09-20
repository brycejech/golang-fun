package userController

import (
	userService "go-api/services/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users := userService.GetUsers()

	c.JSON(200, users)
}

type newUser struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func CreateUser(c *gin.Context) {
	var params newUser
	err := c.ShouldBind(&params)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad Request")
	} else {

		user, err := userService.CreateUser(params.FirstName, params.LastName, params.Email, params.Password)
		if err != nil {
			c.String(http.StatusInternalServerError, "Server Error")
			return
		}

		c.JSON(http.StatusOK, user)
	}

}
