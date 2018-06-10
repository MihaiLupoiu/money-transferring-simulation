package util

import (
	"github.com/MihaiLupoiu/money-transferring-simulation/models"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func InitDb() *gorm.DB {
	// Openning file
	db, err := gorm.Open("sqlite3", "./data.db")
	// Display SQL queries
	db.LogMode(true)

	// Error
	if err != nil {
		panic(err)
	}
	// Creating the table
	if !db.HasTable(&models.Users{}) {
		db.CreateTable(&models.Users{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&models.Users{})
	}

	if !db.HasTable(&models.Transfer{}) {
		db.CreateTable(&models.Transfer{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&models.Transfer{})
	}

	return db
}