package pages

import (
	"fmt"
	"koda-b8-db5/utils"
)

func Change() {
	utils.ListContacts()

	for {
		var id int
		var change string
		var data string
		for {
			var input int
			fmt.Printf("\nInput number ID want change : ")
			_, err1 := fmt.Scanln(&input)

			if err1 != nil {
				continue
			}
			id = input
			break
		}
		for {
			var input int
			fmt.Printf("\n\n1. First Name")
			fmt.Printf("\n2. Last Name")
			fmt.Printf("\n3. Email")
			fmt.Printf("\n4. Phone")
			fmt.Printf("\n\n\nWhat you whant change : ")
			_, err2 := fmt.Scanln(&input)

			if err2 != nil {
				continue
			}
			switch input {
			case 1:
				change = "First_name"
			case 2:
				change = "Last_name"
			case 3:
				change = "Email"
			case 4:
				change = "Phone"
			default:
				continue
			}
			break
		}

		for {
			var input string
			fmt.Printf("\nNew %s : ", change)
			_, err3 := fmt.Scanln(&input)

			if err3 != nil {
				fmt.Printf("Dont have sapcing and not empty")
				var dump string
				fmt.Scanln(&dump)
				continue
			}
			data = input
			break
		}
		update := utils.UpdateContact(change, id, data)
		if update == true {
			fmt.Printf("Succes Change")
			fmt.Scanln()
			return
		}
	}
}
