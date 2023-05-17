package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Companies struct {
	Name    string `gorm:"primary_key" json:"name"`
	Created int    `json:"created"`
	Product string `json:"product"`
}

func DBConnection() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test2.db"), &gorm.Config{})

	if err != nil {
		return db, err
	}

	err = db.AutoMigrate(&Companies{})

	if err != nil {
		return nil, err
	}

	return db, nil

}
