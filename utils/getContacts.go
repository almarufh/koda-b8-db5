package utils

import (
	"context"
	"fmt"
	"koda-b8-db5/database"
	"sync"

	"github.com/jackc/pgx/v5"
)

func GetContacts(wg *sync.WaitGroup, signal *chan string) {
	conn, _ := Conn()
	defer conn.Close(context.Background())
	rows, err := conn.Query(context.Background(), `
		SELECT * FROM "contacts";
	`)

	if err != nil {
		fmt.Println("Gagal memuat table")
	}

	contacts, err := pgx.CollectRows(rows, pgx.RowToStructByName[database.Contact])

	if err != nil {
		fmt.Println("Gagal collect Table")
	}
	newContacts := database.Contacts()
	*newContacts = contacts
	*signal <- "success"
	defer wg.Done()
}
