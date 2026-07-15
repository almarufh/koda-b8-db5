package utils

import (
	"context"
)

func Post(query string, id int, data string) bool {
	conn, _ := Conn()
	defer conn.Close(context.Background())
	if data == "delete" {
		_, err := conn.Query(context.Background(), query, id)

		if err != nil {
			return false
		} else {
			return true
		}
	}
	_, err := conn.Query(context.Background(), query, id, data)

	if err != nil {
		return false
	} else {
		return true
	}
}
