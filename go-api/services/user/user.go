package userService

import (
	"encoding/json"
	fileService "go-api/services/file"

	"github.com/google/uuid"
)

type User struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func GetUsers() []User {
	fileContents := fileService.GetFileContents("data/users.json")

	var users = []User{}
	json.Unmarshal([]byte(fileContents), &users)

	return users
}

func CreateUser(firstName string, lastName string, email string, password string) (newUser User, err error) {
	user := User{
		Id:        uuid.New().String(),
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
	}

	users := GetUsers()

	users = append(users, user)

	bytes, err := json.Marshal(users)

	if err != nil {
		return User{}, err
	}

	fileService.WriteFile("data/users.json", bytes)

	return user, nil
}
