package model

import (
	"database/sql"
	"fmt"
	"log"
)

var con *sql.DB

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:016825353#zx@tcp(127.0.0.1:3306)/test_go")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("db connected...")
	con = db
	return db
}
