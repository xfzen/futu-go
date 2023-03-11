package database

import (
	"futuq/config"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

var db *gorm.DB

func SetupDatabaseV2(c config.Config) {
	driver := c.Database.Driver
	source := c.Database.Source

	logx.Infof("Database driver: %v", driver)
	tmpdb, err := setupDatabase(driver, source)
	if err != nil {
		panic("failed to connect database")
	}

	db = tmpdb

	// 数据库迁移
	err = AutoMigrate()
	if err != nil {
		panic(err)
	}
}

func SetupDatabase() {
	c := config.GetConf()
	driver := c.Database.Driver
	source := c.Database.Source

	logx.Infof("Database driver: %v, source: %v", driver, source)
	tmpdb, err := setupDatabase(driver, source)
	if err != nil {
		panic("failed to connect database")
	}

	db = tmpdb

	// 数据库迁移
	err = AutoMigrate()
	if err != nil {
		panic(err)
	}
}
