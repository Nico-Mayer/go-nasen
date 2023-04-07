package models

import (
	"errors"
	"fmt"

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
		return fmt.Errorf("error inserting nase: %w", err)
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

func GetNasen(id string) ([]Nase, error) {
	query := "SELECT id, userid, authorid, reason FROM nasen WHERE userid = $1"
	rows, err := db.DB.Query(query, id)
	if err != nil {
		return nil, fmt.Errorf("error getting nasen: %w", err)
	}

	var nasen []Nase

	defer rows.Close()

	for rows.Next() {
		var nase Nase
		err := rows.Scan(&nase.ID, &nase.UserID, &nase.AuthorID, &nase.Reason)
		if err != nil {
			return nil, fmt.Errorf("error scanning nase row: %w", err)
		}
		nasen = append(nasen, nase)
	}

	return nasen, nil
}

func CountNasen(id string) (int, error) {
	query := "SELECT COUNT(*) FROM nasen WHERE userid = $1"
	var count int
	err := db.DB.QueryRow(query, id).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("error counting nasen: %w", err)
	}
	return count, nil
}
