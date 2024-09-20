package database

import (
	"user-api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	dsn := "root:root@tcp(localhost:3306)/management"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return err
}

func Migrate() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		panic("Failed to migrate database: " + err.Error())
	}
}
