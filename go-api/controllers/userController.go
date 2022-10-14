package controllers

import (
	"go-api/entity"
	"go-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetUsers(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
	GetUser(ctx *gin.Context)
}

type userController struct{}

var (
	userService = services.NewUserService()
)

func NewUserController() UserController {
	return &userController{}
}

func (controller *userController) GetUsers(c *gin.Context) {
	users := userService.GetUsers()

	c.JSON(200, users)
}

func (controller *userController) CreateUser(c *gin.Context) {
	var params entity.NewUser
	err := c.ShouldBind(&params)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad Request")
	} else {

		user, err := userService.CreateUser(params.FirstName, params.LastName, params.Email, params.Username, params.Password)
		if err != nil {
			c.String(http.StatusInternalServerError, "Server Error")
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

func (controller *userController) GetUser(c *gin.Context) {
	id := c.Param("id")

	if user, ok := userService.GetUser(id); ok {
		c.JSON(http.StatusOK, user)
		return
	}

	c.String(http.StatusNotFound, "Not Found")
}
