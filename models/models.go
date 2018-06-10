package models

import "time"

type Users struct {
	Id        int     `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Firstname string  `gorm:"not null" form:"firstname" json:"firstname"`
	Lastname  string  `gorm:"not null" form:"lastname" json:"lastname"`
	Mail      string  `gorm:"not null" form:"mail" json:"mail"`
	Phone     string  `gorm:"" form:"phone" json:"phone"`
	Balance   float64 `gorm:"" form:"balance" json:"balance"`
	Currency  string  `gorm:"not null;default:Euro" form:"currency" json:"currency"`
}

type Deposit struct {
	Id     int       `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Amount float64   `gorm:"" form:"amount" json:"amount"`
	Date   time.Time `gorm:"" form:"date" json:"date"`
}

type Transfer struct {
	Id       int       `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	SenderId int       `gorm:"not null" form:"senderid" json:"senderid"`
	Amount   float64   `gorm:"" form:"amount" json:"amount"`
	Date     time.Time `gorm:"" form:"date" json:"date"`
}
