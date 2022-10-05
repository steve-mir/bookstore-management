package models

import (
	"github.com/jinzhu/gorm"
	"github.com/steve-mir/bookstore-management/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	name        string `gorm:""json:"name"`
	author      string `json:"author"`
	publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}
