package models

import (
	"futuq/database"

	"gorm.io/gorm"
)

func GetDB() *gorm.DB {
	return database.GetDB()
}
