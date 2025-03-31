package database

import (
	"log"

	envLoader "database.learning/start/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func ConnectDB() (*gorm.DB, error) {

	if db != nil {
		return db, nil
	}

	dsn := envLoader.LoadEnv("DB_DSN")
	if dsn == "" {
		log.Fatal("DB_DSN environment variable is not set")
	}

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	log.Print("Connected to the database")
	return db, nil
}
