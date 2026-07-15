package utils

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func UpdateContact(conn *pgx.Conn, first string, last string, email string, phone string) {
	res, err := conn.Query(context.Background(), `
		UPDATE "contacts" SET "email" 
	`, first, last, email, phone)

	if err != nil {
		fmt.Println("Gagal membuat contact")
	} else {
		fmt.Println("Create contact Succes")
	}

	defer res.Close()
}
