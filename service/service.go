package service

import (
	"laporpak/model"
)

var users []model.User
var lastID = 0

func GetAllUsers() []model.User {
	return users
}

func CreateUser(name string) model.User {
	lastID++

	user := model.User{
		ID:   lastID,
		Name: name,
	}

	users = append(users, user)
	return user
}
