package akmodels

import "futuq/database/model"

type TModelYearlyInfo struct {
	model.Model

	TYearlyInfo

	model.ControlBy
	model.ModelTime
}

type TYearlyInfo struct {
	Symbol string `json:"symbol" gorm:"column:symbol;comment:symbol"`
	Year   int    `json:"year" gorm:"column:year;comment:年份"`

	Jan float64 `json:"m1" gorm:"column:m1;comment:1月升跌幅"`
	Feb float64 `json:"m2" gorm:"column:m2;comment:2月升跌幅"`
	Mar float64 `json:"m3" gorm:"column:m3;comment:3月升跌幅"`
	Apr float64 `json:"m4" gorm:"column:m4;comment:4月升跌幅"`
	May float64 `json:"m5" gorm:"column:m5;comment:5月升跌幅"`
	Jun float64 `json:"m6" gorm:"column:m6;comment:6月升跌幅"`
	Jul float64 `json:"m7" gorm:"column:m7;comment:7月升跌幅"`
	Aug float64 `json:"m8" gorm:"column:m8;comment:8月升跌幅"`
	Sep float64 `json:"m9" gorm:"column:m9;comment:9月升跌幅"`
	Oct float64 `json:"m10" gorm:"column:m10;comment:10月升跌幅"`
	Nov float64 `json:"m11" gorm:"column:m11;comment:11月升跌幅"`
	Dec float64 `json:"m12" gorm:"column:m12;comment:12月升跌幅"`
}

func (TModelYearlyInfo) TableName() string {
	return "t_yearly_change"
}

func (e *TModelYearlyInfo) Generate() model.ActiveRecord {
	o := *e
	return &o
}

func (e *TModelYearlyInfo) GetId() interface{} {
	return e.Id
}
