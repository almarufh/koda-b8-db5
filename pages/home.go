package pages

import (
	"fmt"
	"koda-b8-db5/utils"
	"strings"

	"github.com/jackc/pgx/v5"
)

func Home(conn *pgx.Conn) {
	for {
		var input string
		fmt.Printf("Management Contacts\n\n")
		fmt.Printf("1. List All Contacts\n")
		fmt.Printf("2. Add Contact\n")
		fmt.Printf("3. Change Contact\n")
		fmt.Printf("4. Delete Contact\n")
		fmt.Printf("\nChose a menu :  ")
		fmt.Scanf("%s", &input)

		switch strings.ToLower(input) {
		case "1":
			fmt.Printf("Dalam pengembangan\n\n")
			utils.ListCOntacts(conn)
		case "2":
			fmt.Printf("Dalam pengembangan")
		case "3":
			fmt.Printf("Dalam pengembangan")
		case "4":
			fmt.Printf("Dalam pengembangan")
		default:
			continue
		}
	}

}
