package database

import "gorm.io/gorm"

type DB_Type struct {
	*gorm.DB
}

var DB *DB_Type
