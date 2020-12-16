package model

import "log"

func CreateTodo(name, todo string) error {
	insert, error := con.Query("INSERT INTO TODO (name,todo) VALUES(?,?)", name, todo)
	defer insert.Close()
	if error != nil {
		log.Fatal(error)
	}
	return nil

}
