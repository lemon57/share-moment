package main

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	db, err := sql.Open("pgx", "postgres://bjj:sharemoment@localhost:5432/sharemoment?sslmode=disable")
	//db, err := sql.Open("pgx", "host=localhost port=5432 user=bjj password=sharemoment dbname=sharemoment sslmode=disable")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected!")

	defer db.Close()
}
