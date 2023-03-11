package database

import "futuq/domain/models/histdata"

// 数据库迁移
func AutoMigrate() error {
	err := db.AutoMigrate(
		new(histdata.TIndexMonthly),
	)

	return err
}
