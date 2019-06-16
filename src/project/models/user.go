package models

import (
	"github.com/jinzhu/gorm"
)

// User represent table
type User struct {
	gorm.Model
	Name     string `gorm:"size:255;NOT NULL"`
	Email    string `gorm:"size:100;NOT NULL"`
	Password string `gorm:"size:255;NOT NULL"`
	Role     string `gorm:"size:64;DEFAULT:'buyer'"`
}

// TableName is gorm function to name the table
func (User) TableName() string {
	return "users"
}
