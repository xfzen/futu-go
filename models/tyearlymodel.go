package models

import (
	"futuq/database"
	"futuq/pkg/utils"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type TYearlyModel struct {
	db *gorm.DB
}

func NewTYealModel() *TYearlyModel {
	return &TYearlyModel{
		db: database.GetDB(),
	}
}

// Batch Insert
func (m *TYearlyModel) BatchInsert(dataList []TYearlyData) error {
	for _, v := range dataList {
		m.Insert(&v)
	}

	return nil
}

// Insert
func (m *TYearlyModel) Insert(data *TYearlyData) error {
	// 通过uid查询，存在则更新；不存在则新增
	if m.db.Model(TYearlyData{}).Where("symbol = ? and year = ?", data.Symbol, data.Year).Updates(&data).RowsAffected == 0 {
		m.db.Create(&data)
	}

	return nil
}

func (m *TYearlyModel) List(cond map[string]string, page int, size int) (*[]TYearlyData, int64, error) {
	var recordArr []TYearlyData

	//  根据时间范围进行查找
	startDate, startExist := cond["startDate"]
	if startExist {
		startDate = cond["startDate"] + " 00:00:00"
		delete(cond, "startDate")
	}

	endDate, endExist := cond["endDate"]
	if endExist {
		endDate = cond["endDate"] + " 23:59:59"
		delete(cond, "endDate")
	}

	var dateCond string
	if startExist && endExist {
		dateCond = " and createdAt >= '" + startDate + "' and createdAt <= '" + endDate + "'"
	}

	//  将map转成sql条件
	sqlCond := utils.ListMap2Str(cond) + dateCond

	//  分页
	offset := (page - 1) * size
	if offset < 0 {
		offset = 0
	}

	var count int64
	err := m.db.Where(sqlCond).Order("createdAt desc").Limit(size).Offset(offset).Find(&recordArr).Error
	if err != nil {
		logx.Errorf("List by %s error, reason: %s", TYearlyData{}.TableName(), err)
		return &recordArr, 0, err
	}

	//  单独计数，count和limit、offset不能混用
	sqlCond = sqlCond + " and deletedAt IS NULL"
	m.db.Table(TYearlyData{}.TableName()).Where(sqlCond).Count(&count)

	return &recordArr, count, nil
}
