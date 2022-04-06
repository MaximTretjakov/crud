package main

import (
	"database/sql"
	"fmt"
	_ "github/lib/pq"
	"log"
)

func main() {
	var id int
	connStr := "user=postgres password=postgres dbname=postgres sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	result, err := db.Query("select id, name from users where id = ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	result.Scan(&id)
	fmt.Println(id)
}
