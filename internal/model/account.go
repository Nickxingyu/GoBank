package model

import (
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
	if err := db.Find(&account, accountID).Error; err != nil {
		return nil, err
	}
	return &account, nil
}

func InsertAccount(account Account) (*Account, error) {
	if err := db.Create(&account).Error; err != nil {
		return nil, err
	}

	return &account, nil
}
