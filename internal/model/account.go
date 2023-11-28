package model

import "gorm.io/gorm"

type Account struct {
	gorm.Model

	UserID   uint
	CoinType string
	Balance  float64
}
