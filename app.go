package main

import (
	"github.com/Nickxingyu/GoBank/internal/app"
	"github.com/Nickxingyu/GoBank/internal/database"
)

func main() {
	db := database.Postgres{
		Host:     "localhost",
		User:     "postgres",
		Password: "admin",
		DBname:   "GoBank",
		Port:     "8080",
		SSLmode:  "disable",
	}
	db.Conn()

	server := app.NewApiServer(":3000")
	server.Run()
}
