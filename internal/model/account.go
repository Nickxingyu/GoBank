package model

import (
	"fmt"

	"github.com/Nickxingyu/GoBank/internal/database"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model

	UserID   uint
	CoinType string
	Balance  float64
}

func init() {
	fmt.Println("account model init")
	database.DB.AutoMigrate(&Account{})
}

func FindAccountByID(accountID uint) (*Account, error) {
	account := Account{}
	if err := database.DB.First(&account, accountID).Error; err != nil {
		return nil, err
	}
	return &account, nil
}

func InsertAccount(account Account) (*Account, error) {
	if err := database.DB.Create(&account).Error; err != nil {
		return nil, err
	}

	return &account, nil
}

func SaveAccountByID(accountID uint, account Account) (*Account, error) {
	account.ID = accountID
	if err := database.DB.Save(&account).Error; err != nil {
		return nil, err
	}

	return &account, nil
}

func DeleteAccountByID(accountID uint) (*Account, error) {
	var account Account

	err := database.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&account, accountID).Error; err != nil {
			return err
		}
		if err := tx.Delete(&account).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &account, nil
}
