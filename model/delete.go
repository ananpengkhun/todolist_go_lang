package model

import (
	"log"
)

func DeleteTodo(id int) error {
	delete, err := con.Query("DELETE FROM TODO WHERE todo_id = ?", id)
	defer delete.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	return nil
}
