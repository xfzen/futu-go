package database

import (
	"database/sql"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB_FILE_SQLITE = "futuq.db"

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
	createDatabaseNotExist(source)

	db, err := gorm.Open(mysql.Open(source), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db, nil
}

// 若数据库不存在则自动创建
// root:root123456@tcp(127.0.0.1:3306)/futuq?charset=utf8mb4&parseTime=True&loc=Local&timeout=1000ms
func createDatabaseNotExist(source string) error {
	// 找到字符串"/"和"?"的位置
	index1 := strings.Index(source, "/")
	index2 := strings.Index(source, "?")

	sourceRoot := source[:index1] + "/"
	dbName := source[index1+1 : index2]

	// 创建数据库连接
	db, err := sql.Open("mysql", sourceRoot)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	if err != nil {
		logx.Infof("%v", err)
		return nil
	}

	return nil
}
