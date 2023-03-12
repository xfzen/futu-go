package akfeeds

import (
	"encoding/json"

	"github.com/zeromicro/go-zero/core/logx"
)

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
