package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB_FILE_SQLITE = "safetytools.db"

func setupDatabase(driver, source string) (*gorm.DB, error) {
	switch driver {
	case "sqlite":
		return initSqlite(DB_FILE_SQLITE)

	case "mysql":
		return initMysql(source)

	default:
		return initSqlite(DB_FILE_SQLITE)
	}
}

func initSqlite(dbfile string) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(DB_FILE_SQLITE), &gorm.Config{PrepareStmt: true})
}

func initMysql(source string) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(source), &gorm.Config{})
}
