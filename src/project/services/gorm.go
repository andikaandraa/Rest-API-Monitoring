package services

import (
	"github.com/jinzhu/gorm"
)

// ObjDB represent DB object
type ObjDB struct {
	DB *gorm.DB
}
