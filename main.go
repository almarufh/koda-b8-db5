package main

import (
	"fmt"
	"koda-b8-db5/pages"
	"koda-b8-db5/utils"
)

func main() {
	conn, err := utils.Conn()
	if err != nil {
		fmt.Println("Gagal membuat conection ke database")
	}

	pages.Home(conn)

	// utils.CreateContact(conn, "Siti", "Zulaikha", "zulaikha529@gmail.com", "085399376716")

}
