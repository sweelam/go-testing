package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func connect() mysql.Config {
	fmt.Println("Test")

	return mysql.Config{
		User:   "sweelam",
		Passwd: "sweelam_224466",
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "sweprofile",
	}
}

func Init() {
	cfg := connect()
	var err error

	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	println("Connected to the database")
}
