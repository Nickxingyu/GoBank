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

func FindUserByID(userID uint) (*User, error) {
	user := User{}
	if err := database.DB.First(&user, userID).Error; err != nil {
		return nil, err
	}

	return &user, nil
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

func SaveUserByID(userID uint, user User) (*User, error) {
	user.ID = userID
	if err := database.DB.Save(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func DeleteUserByID(userID uint) (*User, error) {
	user := User{}
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&user, userID).Error; err != nil {
			return err
		}

		if err := tx.Delete(&user).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &user, nil
}
