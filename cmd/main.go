package main

import (
	envLoader "database.learning/start/config"
	database "database.learning/start/internal/database/gorm"
)

func main() {
	envLoader.LoadAllEnv()
	database.ConnectDB()
	StartServer()
}
