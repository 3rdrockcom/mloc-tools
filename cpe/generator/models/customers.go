package models

type Customers struct {
	ID           uint `gorm:"PRIMARY_KEY"`
	Email        string
	Gender       string
	FirstName    string
	LastName     string
	MobileNumber string
}
