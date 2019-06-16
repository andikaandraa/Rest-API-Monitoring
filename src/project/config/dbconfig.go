package config

import (
	"project/models"

	"github.com/jinzhu/gorm"
)

// DBInit is initialing DB connection
func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("Failed to connect")
	}

	db.AutoMigrate(models.User{})
	return db
}
