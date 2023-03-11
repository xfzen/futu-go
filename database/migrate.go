package database

// 数据库迁移
func AutoMigrate() error {
	err := db.AutoMigrate()

	return err
}
