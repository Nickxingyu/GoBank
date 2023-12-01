package model

import (
	"github.com/Nickxingyu/GoBank/internal/database"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model

	UserID   uint
	CoinType string
	Balance  float64
}

func FindAccountById(accountID uint) (*Account, error) {
	account := Account{}
	if err := database.DB.Find(&account, accountID).Error; err != nil {
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
	if err := database.DB.Find(&account, accountID).Error; err != nil {
		return nil, err
	}
	if err := database.DB.Delete(&account).Error; err != nil {
		return nil, err
	}

	return &account, nil
}
