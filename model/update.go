package model

import (
	"log"
)

func UpdateTodo(status, id int) error {
	update, err := con.Query("UPDATE TODO SET status = ? WHERE todo_id = ? ", status, id)
	defer update.Close()
	if err != nil {
		log.Fatal(err.Error())
	}

	return nil
}
