package models

import (
	"futuq/database"
	"futuq/pkg/utils"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type TMonthlyIndexModel struct {
	db *gorm.DB
}

func NewTMonthlyIndexModel() *TMonthlyIndexModel {
	return &TMonthlyIndexModel{
		db: database.GetDB(),
	}
}

// Batch Insert
func (m *TMonthlyIndexModel) BatchInsert(dataList []TMonthlyIndexData) error {
	for _, v := range dataList {
		m.Insert(&v)
	}

	return nil
}

// Insert
func (m *TMonthlyIndexModel) Insert(data *TMonthlyIndexData) error {
	// 通过uid查询，存在则更新；不存在则新增
	if m.db.Model(TMonthlyIndexData{}).Where("date = ?", data.Date).Updates(&data).RowsAffected == 0 {
		m.db.Create(&data)
	}

	return nil
}

func (m *TMonthlyIndexModel) List(cond map[string]string, page int, size int) (*[]TMonthlyIndexData, int64, error) {
	var recordArr []TMonthlyIndexData

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
		logx.Errorf("List by %s error, reason: %s", TMonthlyIndexData{}.TableName(), err)
		return &recordArr, 0, err
	}

	//  单独计数，count和limit、offset不能混用
	sqlCond = sqlCond + " and deletedAt IS NULL"
	m.db.Table(TMonthlyIndexData{}.TableName()).Where(sqlCond).Count(&count)

	return &recordArr, count, nil
}
