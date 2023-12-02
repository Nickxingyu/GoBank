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

type EmailHasBeenUsedError struct {
	Email string
}

func (e EmailHasBeenUsedError) Error() string {
	return "This email " + e.Email + " already be used."
}

func InsertUser(user User) (*User, error) {
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		userModel := User{}
		if err := tx.Where(&User{Email: user.Email}).First(&userModel).Error; err != nil {
			if err != gorm.ErrRecordNotFound {
				return err
			}

			if err := tx.Create(&user).Error; err != nil {
				return err
			}

			return nil
		}

		return EmailHasBeenUsedError{Email: user.Email}
	})
	if err != nil {
		return nil, err
	}

	return &user, nil
}
