package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	UserID   string `json:"userId" gorm:"primaryKey"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"index"`
	CPF      string `json:"cpf" gorm:"index"`
	Password string `json:"password"`
}

func (u *User) BeforeCreate(scope *gorm.DB) (err error) {
	u.UserID = uuid.NewString()

	return
}
