package utils

import (
	"fmt"
	"koda-b8-db5/database"
	"sync"

	"github.com/jackc/pgx/v5"
)

func ListCOntacts(conn *pgx.Conn) {
	wg := sync.WaitGroup{}
	signal := make(chan string)
	wg.Add(1)
	go GetContacts(conn, &wg, &signal)

	i := ""
	for i != "success" {
		if <-signal == "success" {
			i = "success"
		}
	}

	wg.Wait()

	contacts := database.Contacts()

	if len(*contacts) < 1 {
		fmt.Printf("Contacts empty\n\nPress enter to back! ")
		fmt.Scanf("\n")
		return
	}

	fmt.Printf("---[ List All COntacts ]---\n\n")

	for _, res := range *contacts {
		fmt.Printf("   -----------------------\n")
		fmt.Printf("ID [%d]\n", res.Id)
		fmt.Printf("   Name  : %s %s\n", res.First_name, res.Last_name)
		fmt.Printf("   Email : %s\n", res.Email)
		fmt.Printf("   Phone : %s\n", res.Phone)
		fmt.Printf("   -----------------------\n\n")
	}

	fmt.Printf("\n\nPress enter to back!  ")
	fmt.Scanf("\n")
}
