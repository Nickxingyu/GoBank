package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	Host     string
	User     string
	Password string
	DBname   string
	Port     string
	SSLmode  string
}

func (p Postgres) getDsn() string {
	return "host=" + p.Host + " user=" + p.User + " password=" + p.Password + " dbname=" + p.DBname + " port=" + p.Port + " sslmode=" + p.SSLmode
}

func (p Postgres) Connect() {
	dsn := p.getDsn()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	DB = db
}
