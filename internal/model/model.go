package model

import (
	"github.com/Nickxingyu/GoBank/internal/database"
)

var db *database.DB_Type

func ConnectDatabase() {
	database.Postgres{
		Host:     "localhost",
		User:     "postgres",
		Password: "admin",
		DBname:   "GoBank",
		Port:     "8080",
		SSLmode:  "disable",
	}.Connect()
	db = database.DB

	db.AutoMigrate(&Account{})
	db.AutoMigrate(&User{})
}
