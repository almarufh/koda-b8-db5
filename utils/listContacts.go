package utils

import (
	"fmt"
	"koda-b8-db5/database"

	"github.com/jackc/pgx/v5"
)

func ListCOntacts(conn *pgx.Conn) {
	GetContacts(conn)

	for _, res := range *(database.Contacts()) {
		fmt.Printf("ID    : %d\n", res.Id)
		fmt.Printf("Name  : %s %s\n", res.First_name, res.Last_name)
		fmt.Printf("Email : %s\n", res.Email)
		fmt.Printf("Phone : %s\n", res.Phone)
		fmt.Println()
	}
}
