package model

import (
	"github.com/Nickxingyu/GoBank/internal/database"
)

func ConnectDatabase() {
	database.Postgres{
		Host:     "localhost",
		User:     "postgres",
		Password: "admin",
		DBname:   "GoBank",
		Port:     "8080",
		SSLmode:  "disable",
	}.Connect()

	database.DB.AutoMigrate(&Account{})
	database.DB.AutoMigrate(&User{})
}
