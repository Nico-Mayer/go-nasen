package models

import (
	"errors"

	"github.com/google/uuid"
	"github.com/nico-mayer/go-api/db"
)

type Nase struct {
	ID       uuid.UUID `json:"id"`
	UserID   string    `json:"userid"`
	AuthorID string    `json:"authorid"`
	Reason   string    `json:"reason"`
}

func InsertNase(nase *Nase) error {
	query := "INSERT INTO nasen (id, userid, authorid, reason) VALUES ($1, $2, $3, $4)"
	_, err := db.DB.Exec(query, nase.ID, nase.UserID, nase.AuthorID, nase.Reason)
	if err != nil {
		panic(err)
	}

	return nil
}

func ValidateNase(nase *Nase) error {
	if nase.ID == uuid.Nil {
		return errors.New("missing or invalid ID")
	}

	if nase.UserID == "" {
		return errors.New("missing or invalid UserID")
	}

	if nase.AuthorID == "" {
		return errors.New("missing or invalid AuthorID")
	}

	return nil
}
