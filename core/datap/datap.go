package datap

import (
	"encoding/json"

	"github.com/zeromicro/go-zero/core/logx"
)

// 处理指数实时数据
func HandleIndexSpot(jsonstr string) error {
	var indexSpot []map[string]interface{}

	err := json.Unmarshal([]byte(jsonstr), &indexSpot)
	if err != nil {
		logx.Errorf("json.Unmarshal err: %v", err)
		return err
	}

	for k, v := range indexSpot {
		logx.Infof("k: %v, v: %v", k, v)
	}

	return nil
}

// 处理指数历史行情
func HandleIndexData(symbol, jsonstr string) error {
	var indexDatas []map[string]interface{}

	err := json.Unmarshal([]byte(jsonstr), &indexDatas)
	if err != nil {
		logx.Errorf("json.Unmarshal err: %v", err)
		return err
	}

	// convert to monthly data
	monthlyDataList := ConvertToMonthlyData(symbol, indexDatas)

	// 年度变化数据处理与汇总
	yearlyDataList := ProcessMonthlyData(monthlyDataList)
	logx.Infof("yearlyDataList: size: %v", len(yearlyDataList))

	// insert monthlyData to db
	HandleMonthlyData(monthlyDataList)

	// insert yearlyList to db
	YearlyDataHandle(yearlyDataList)

	return nil
}
