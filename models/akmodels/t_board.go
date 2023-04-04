package akmodels

import "futuq/database/model"

type TBoardInfo struct {
	Symbol string  `json:"symbol" gorm:"column:symbol;comment:代码"`
	Name   string  `json:"name" gorm:"column:year;comment:股票名字"`
	Market string  `json:"mkt" gorm:"column:mkt;comment:市场类型"`
	Close  float64 `json:"close" gorm:"column:close;comment:收盘价"`
	High   float64 `json:"high" gorm:"column:high;comment:最高价"`
	Low    float64 `json:"low" gorm:"column:close;comment:最低价"`
}

type TModelBoardInfo struct {
	model.Model
	TBoardInfo
	model.ControlBy
	model.ModelTime
}

func (TModelBoardInfo) TableName() string {
	return "t_board"
}

func (e *TModelBoardInfo) Generate() model.ActiveRecord {
	o := *e
	return &o
}

func (e *TModelBoardInfo) GetId() interface{} {
	return e.Id
}
