package utils

import (
	"context"
	"fmt"
	"koda-b8-db5/database"

	"github.com/jackc/pgx/v5"
)

func GetContacts(conn *pgx.Conn) {
	rows, err := conn.Query(context.Background(), `
		SELECT * FROM "contacts";
	`)

	if err != nil {
		fmt.Println("Gagal memuat table")
	}

	defer rows.Close()

	contacts, err := pgx.CollectRows(rows, pgx.RowToStructByName[database.Contact])

	if err != nil {
		fmt.Println("Gagal collect Table")
	}
	newContacts := database.Contacts()
	*newContacts = contacts
}
