package pages

import (
	"fmt"
	"koda-b8-db5/database"
	"koda-b8-db5/utils"
)

func AddContact() {
	utils.Clear()
	fmt.Printf("--[ ADD CONCTACT ]---\n\n")
	var first string
	var last string
	var email string
	var phone string
	fmt.Printf("\nFirst Name : ")
	fmt.Scanf("%s", &first)
	fmt.Printf("Last Name  : ")
	fmt.Scanf("%s", &last)
	fmt.Printf("Email      : ")
	fmt.Scanf("%s", &email)
	fmt.Printf("Phone      : ")
	fmt.Scanf("%s", &phone)
	newContact := database.Contact{
		First_name: first,
		Last_name:  last,
		Email:      email,
		Phone:      phone,
	}

	res := utils.CreateContact(newContact)
	if res == true {
		fmt.Printf("\nAdd contact %s %s success\n\n", first, last)
		fmt.Printf("Press Enter to back! ")
		fmt.Scanf("\n")
	}
}
