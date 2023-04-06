package models

import (
	"errors"

	"github.com/nico-mayer/go-api/db"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func InsertUser(user User) error {
	_, err := db.DB.Exec("INSERT INTO users (id, name) VALUES ($1, $2)", user.ID, user.Name)
	if err != nil {
		panic(err)
	}
	return nil
}

func GetUser(id string) User {
	var user User
	err := db.DB.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name)
	if err != nil {
		panic(err)
	}
	return user
}

func GetAllUser() []User {
	var users []User
	rows, err := db.DB.Query("SELECT * FROM users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return users
}

func ValidateUser(user *User) error {
	if user.ID == "" {
		return errors.New("missing or invalid ID")
	}

	if user.Name == "" {
		return errors.New("missing or invalid Name")
	}

	return nil
}
