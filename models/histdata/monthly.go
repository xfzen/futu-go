package histdata

import "futuq/database/model"

type ModelIndexMonthly struct {
	Symbol string  `json:"symbol" gorm:"column:symbol;comment:symbol"`
	Date   string  `json:"date" gorm:"column:date;comment:年月（如2023-03）"`
	Year   int     `json:"year" gorm:"column:year;comment:年份"`
	Month  int     `json:"month" gorm:"column:month;comment:月份"`
	Open   float64 `json:"open" gorm:"column:open;comment:开盘价"`
	Close  float64 `json:"close" gorm:"column:close;comment:收盘价"`
	High   float64 `json:"high" gorm:"column:high;comment:最高价"`
	Low    float64 `json:"low" gorm:"column:low;comment:最低价"`
	Change float64 `json:"change" gorm:"column:change;comment:升跌幅度"`
}

type TIndexMonthly struct {
	model.Model

	ModelIndexMonthly

	model.ControlBy
	model.ModelTime
}

func (TIndexMonthly) TableName() string {
	return "t_index_monthly"
}

func (e *TIndexMonthly) Generate() model.ActiveRecord {
	o := *e
	return &o
}

func (e *TIndexMonthly) GetId() interface{} {
	return e.Id
}
