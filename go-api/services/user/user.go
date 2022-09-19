package userService

import (
	"encoding/json"
	fileService "go-api/services/file"
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
