package database

import (
	"database/sql"
	"log"
)

func Read(qry string) *sql.Rows {
	rs, error := db.Query(qry)

	if error != nil {
		log.Fatal(error)
	}

	return rs
}
