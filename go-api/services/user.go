package services

import (
	"encoding/json"
	"go-api/entity"

	"github.com/google/uuid"
)

type UserService interface {
	GetUsers() []entity.User
	CreateUser(firstName string, lastName string, email string, username string, password string) (entity.User, error)
	GetUser(id string) (user entity.User, ok bool)
}

type userService struct{}

var file FileService = &fileService{}

func NewUserService() UserService {
	return &userService{}
}

func (service *userService) GetUsers() []entity.User {
	fileContents := file.GetFileContents("data/users.json")

	var users = []entity.User{}
	json.Unmarshal([]byte(fileContents), &users)

	return users
}

func (service *userService) CreateUser(firstName string, lastName string, email string, username string, password string) (newUser entity.User, err error) {
	user := entity.User{
		Id:        uuid.New().String(),
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Username:  username,
		Password:  password,
	}

	users := service.GetUsers()

	users = append(users, user)

	bytes, err := json.Marshal(users)

	if err != nil {
		return entity.User{}, err
	}

	file.WriteFile("data/users.json", bytes)

	return user, nil
}

func (service *userService) GetUser(id string) (entity.User, bool) {
	users := service.GetUsers()

	user := entity.User{}
	for _, item := range users {
		if item.Id == id {
			return item, true
		}
	}

	return user, false
}

func (service *userService) UpdateUser(user entity.User) (entity.User, bool) {
	users := service.GetUsers()

	length := len(users)
	for i, u := range users {
		if u.Id == user.Id {
			users = append(users[0:i], user)
			users = append(users, users[i+1:length]...)
			return users, true
		}
	}
}
