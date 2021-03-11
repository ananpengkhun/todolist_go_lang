package model

import "log"

func CreateTodo(name, todo string, status int) error {
	insert, error := con.Query("INSERT INTO TODO (name,todo,status) VALUES(?,?,?)", name, todo, status)
	defer insert.Close()
	if error != nil {
		log.Fatal(error)
	}
	return nil

}
