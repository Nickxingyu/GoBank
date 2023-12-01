package model

import (
	"github.com/Nickxingyu/GoBank/internal/database"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	FirstName string
	LastName  string
	Email     string
	Hash      string
	Accounts  []Account
}

func init() {
	database.DB.AutoMigrate(&User{})
}
