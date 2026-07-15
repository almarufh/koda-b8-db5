package pages

import (
	"fmt"
	"koda-b8-db5/utils"
)

func Delete() {
	var index int
	for {
		for {
			utils.Clear()
			var input int
			utils.ListContacts()
			fmt.Printf("\n\nInput number ID want delete : ")
			_, err := fmt.Scanln(&input)
			if err != nil {
				fmt.Printf("Input wrong")
				continue
			}
			index = input
			break
		}
		query := `DELETE FROM "contacts" WHERE "Id" = $1;`
		res := utils.Post(query, index, "delete")
		if res == true {
			return
		}
		break
	}
}
