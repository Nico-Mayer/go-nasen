package models

import (
	"errors"
	"fmt"

	"github.com/nico-mayer/go-api/db"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Count int    `json:"nasen"`
}

func InsertUser(user *User) error {
	_, err := db.DB.Exec("INSERT INTO users (id, name) VALUES ($1, $2)", user.ID, user.Name)
	if err != nil {
		return fmt.Errorf("error inserting user: %w", err)
	}
	return nil
}

func GetUser(id string) (User, error) {
	var user User
	err := db.DB.QueryRow("SELECT (id, name) FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name)
	if err != nil {
		return User{}, fmt.Errorf("error getting user: %w", err)
	}
	return user, nil
}

func GetAllUsers() ([]User, error) {
	var users []User
	rows, err := db.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, fmt.Errorf("error getting all users: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name)
		if err != nil {
			return nil, fmt.Errorf("error scanning user row: %w", err)
		}
		users = append(users, user)
	}
	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("error with user rows: %w", err)
	}
	return users, nil
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
