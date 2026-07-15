package utils

import (
	"context"
	"koda-b8-db5/database"
)

func CreateContact(data database.Contact) bool {
	conn, _ := Conn()
	_, err := conn.Query(context.Background(), `
		INSERT INTO "contacts" ("First_name", "Last_name", "Email", "Phone") VALUES
		($1, $2, $3, $4)
	`, data.First_name, data.Last_name, data.Email, data.Phone)

	if err != nil {
		return false
	} else {
		return true
	}

}
