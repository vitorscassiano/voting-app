package models

type Authentication struct {
	Email    string `gorm:"primaryKey;index;unique"`
	CPF      string `gorm:"index;unique"`
	Password string
}
