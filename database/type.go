package database

import "time"

type Contact struct {
	Id         int       `db:"Id"`
	First_name string    `db:"First_name"`
	Last_name  string    `db:"Last_name"`
	Email      string    `db:"Email"`
	Phone      string    `db:"Phone"`
	CreatedAt  time.Time `db:"CreatedAt"`
	UpdatedAt  time.Time `db:"UpdatedAt"`
}

var contacts []Contact = []Contact{}

func Contacts() *[]Contact {
	return &contacts
}
