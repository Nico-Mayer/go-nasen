package models

import "errors"

type Person struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
}

var People []Person

func ValidatePerson(p *Person) error {
	if p.ID == "" {
		return errors.New("missing or invalid ID")
	}

	if p.FirstName == "" {
		return errors.New("missing or invalid first name")
	}

	if p.LastName == "" {
		return errors.New("missing or invalid last name")
	}

	if p.Age < 0 {
		return errors.New("invalid age")
	}

	return nil
}

func GetAllPeople() []Person {
	return People
}
