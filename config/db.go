package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" //Needed to setup DB
)

// DB is open pg connection to be used in models
var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("postgres", "postgres://marekparafianowicz:password@localhost/go-server?sslmode=disable")
	if err != nil {
		panic(err)
	}

	if err = DB.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("You connected to your database.")
}
