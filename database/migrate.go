package database

import "futuq/models/akmodels"

// 数据库迁移
func AutoMigrate() error {
	err := db.AutoMigrate(
		new(akmodels.TIndexMonthly),
	)

	return err
}
