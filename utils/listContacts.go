package utils

import (
	"fmt"
	"koda-b8-db5/database"
	"sync"
)

func ListContacts() {
	Clear()
	wg := sync.WaitGroup{}
	signal := make(chan string)
	wg.Add(1)
	go GetContacts(&wg, &signal)

	i := ""
	for i != "success" {
		if <-signal == "success" {
			i = "success"
		}
	}

	wg.Wait()

	contacts := database.Contacts()

	if len(*contacts) < 1 {
		fmt.Printf("Contacts empty\n")
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
}
