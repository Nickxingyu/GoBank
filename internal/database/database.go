package database

import "gorm.io/gorm"

var DB *gorm.DB

func init() {
	Postgres{
		Host:     "localhost",
		User:     "postgres",
		Password: "admin",
		DBname:   "GoBank",
		Port:     "8080",
		SSLmode:  "disable",
	}.Connect()
}
