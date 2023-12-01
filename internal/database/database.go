package database

import (
	"github.com/Nickxingyu/GoBank/internal/config"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	db_config := config.Config.Database.Postgres
	Postgres{
		Host:     db_config.Host,
		User:     db_config.User,
		Password: db_config.Password,
		DBname:   db_config.DBname,
		Port:     db_config.Port,
		SSLmode:  db_config.SSLmode,
	}.Connect()
}
