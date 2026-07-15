package pages

import (
	"fmt"
	"koda-b8-db5/utils"
	"os"
	"strings"

	"github.com/jackc/pgx/v5"
)

func Home(conn *pgx.Conn) {
	for {
		utils.Clear()
		var input string
		fmt.Printf("Management Contacts\n\n")
		fmt.Printf("1. List All Contacts\n")
		fmt.Printf("2. Add Contact\n")
		fmt.Printf("3. Change Contact\n")
		fmt.Printf("4. Delete Contact\n")
		fmt.Printf("\n\n0. Exit\n")
		fmt.Printf("\nChose a menu :  ")
		fmt.Scanf("%s", &input)

		switch strings.ToLower(input) {
		case "1":
			utils.ListContacts()
			fmt.Printf("\n\nPress enter to back!  ")
			fmt.Scanf("\n")
		case "2":
			AddContact()
		case "3":
			Change()
		case "4":
			Delete()
		case "0":
			os.Exit(1)
		default:
			continue
		}
	}

}
