package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	db, err := sql.Open("mysql", "admin:AAAA@tcp(localhost:3309)/db")
	if err != nil {
		fmt.Println("sql open err", err.Error())
		return
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		fmt.Println("mysql with instance err", err.Error())
		return
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://./migration",
		"mysql",
		driver,
	)
	if err != nil {
		fmt.Println("mirate database err", err.Error())
		return
	}

	if err := m.Steps(1); err != nil {

		log.Fatal(err)
	}
}
