package utils

import (
	"context"
	"koda-b8-db5/database"
)

func CreateContact(data database.Contact) bool {
	conn, er := Conn()
	if er != nil {
		return false
	}
	defer conn.Close(context.Background())
	_, err := conn.Exec(context.Background(), `
		INSERT INTO "contacts" ("First_name", "Last_name", "Email", "Phone") VALUES
		($1, $2, $3, $4);
	`, data.First_name, data.Last_name, data.Email, data.Phone)

	if err != nil {
		return false
	}
	return true
}
