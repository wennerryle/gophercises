package main

import (
	"errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func getValue(db *gorm.DB, key string) (string, error) {
	var out AddressBind

	db.Where("path = ?", key).First(&out)

	if out.Path == "" {
		return out.Url, errors.New("no redirect found")
	}

	return out.Url, db.Error
}

func getDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect to db")
	}

	db.AutoMigrate(&AddressBind{})
	db.AllowGlobalUpdate = true

	return db
}

func cleanTable(db *gorm.DB) {
	db.Delete(&AddressBind{})
}

func updateTable(db *gorm.DB, values []AddressBind) {
	db.Create(values)
}
