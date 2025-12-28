package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("hello to you world.")
	connStr := os.Getenv("DATABASE_URL")

	if connStr == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT usesysid, usename FROM pg_catalog.pg_user;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("User ID: %d, Name: %s \n", id, name)
	}
}
