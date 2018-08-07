package main

import (
	"github.com/epointpayment/mloc-tools/cpe/generator/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func NewDB(dbName string) *gorm.DB {
	db, err := gorm.Open("sqlite3", dbName)
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func DoMigrations() {
	// Drop tables
	db.DropTableIfExists(&models.Customers{})
	db.DropTableIfExists(&models.Transactions{})

	// Create schema
	db.AutoMigrate(&models.Customers{})
	db.AutoMigrate(&models.Transactions{})
}
