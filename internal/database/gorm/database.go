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

func ConnectDB() {

	if db != nil {
		return
	}

	dsn := envLoader.LoadEnv("DB_DSN")
	if dsn == "" {
		log.Fatal("DB_DSN environment variable is not set")
		panic("DB_DSN environment variable is not set")
	}

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	log.Print("Connected to the database")
}

func GetDB() *gorm.DB {
	if db == nil {
		log.Fatal("Database connection is not established")
		panic("Database connection is not established")
	}
	return db
}
