package repository

import (
	"laporpak/database"
	"laporpak/model"
)

func CreateUser(name string) (model.User, error) {
	var user model.User

	query := `INSERT INTO users(name) VALUES($1) RETURNING id`

	err := database.DB.QueryRow(query, name).Scan(&user.ID)

	if err != nil {
		return user, err
	}

	user.Name = name
	return user, nil
}

func GetAllUsers() ([]model.User, error) {
	rows, err := database.DB.Query(`SELECT id, name FROM users`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var u model.User
		rows.Scan(&u.ID, &u.Name)
		users = append(users, u)
	}

	return users, nil
}

// func DeleteUser(id INT) ([]model.User, error) {
// 	rows, err := database.DB.Query(`DELETE users where id=($1) FROM users`)
// 	if err != nil {
// 		return nil, err
// 	}

// }
