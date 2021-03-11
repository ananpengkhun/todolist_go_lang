package model

import (
	"log"
)

func UpdateTodo(status, id int, name, todo string) error {
	update, err := con.Query("UPDATE TODO SET status = ?, name = ?, todo = ? WHERE todo_id = ? ", status, name, todo, id)
	defer update.Close()
	if err != nil {
		log.Fatal(err.Error())
	}

	return nil
}
