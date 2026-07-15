package utils

import "fmt"

func UpdateContact(change string, id int, data string) bool {
	var query string
	switch change {
	case "First_name":
		query = `UPDATE "contacts" SET "First_name" = $2 WHERE "Id" = $1;`
	case "Last_name":
		query = `UPDATE "contacts" SET "Last_name" = $2 WHERE "Id" = $1;`
	case "Email":
		fmt.Println("Change Email")
		query = `UPDATE "contacts" SET "Email" = $2 WHERE "Id" = $1;`
	case "Phone":
		query = `UPDATE "contacts" SET "Phone" = $2 WHERE "Id" = $1;`
	default:
		return false
	}

	res := Post(query, id, data)
	return res
}
