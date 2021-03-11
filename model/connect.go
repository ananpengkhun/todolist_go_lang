package model

import (
	"database/sql"
	"fmt"
	"log"
)

var con *sql.DB

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:123456#Qw@tcp(127.0.0.1:3306)/todo_list")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("db connected...")
	con = db
	return db
}
