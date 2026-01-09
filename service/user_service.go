package service

import (
	"laporpak/model"
	"laporpak/repository"
)

var users []model.User
var lastID = 0

func GetAllUsers() ([]model.User, error) {
	return repository.GetAllUsers()
}

func CreateUser(name string) (model.User, error) {
	return repository.CreateUser(name)

}

func UpdateUser(id int, name string) (model.User, bool) {
	for i, user := range users {
		if user.ID == id {
			users[i].Name = name
			return users[i], true
		}
	}

	return model.User{}, false
}

func DeleteUser(id int) bool {
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			return true
		}
	}

	return false
}
