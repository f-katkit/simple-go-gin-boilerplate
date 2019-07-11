package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	db, err = gorm.Open("postgres", "host=0.0.0.0 port=5432 user=gorm dbname=gorm password=gorm sslmode=disable")

	if err != nil {
		panic(err)
	}

	autoMigration()
	InsertSeed()
}

// GetDB is called in models
func GetDB() *gorm.DB {
	return db
}
