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
	//connect to mysql server on localhost:3309
	//on the real server, we should not hard code stuff like this
	db, err := sql.Open("mysql", "admin:AAAA@tcp(localhost:3309)/db")
	if err != nil {
		fmt.Println("sql open err", err.Error())
		return
	}

	//attach mysql client to mysql driver
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		fmt.Println("mysql with instance err", err.Error())
		return
	}

	//load migration from file source(read more about migration sources https://github.com/golang-migrate/migrate#migration-sources)
	//on this project we store migration file on migration folder
	m, err := migrate.NewWithDatabaseInstance(
		"file://./migration",
		"mysql",
		driver,
	)
	if err != nil {
		fmt.Println("mirate database err", err.Error())
		return
	}

	//Do migration to N version
	if err := m.Steps(2); err != nil {
		log.Fatal(err)
		return
	}
}
