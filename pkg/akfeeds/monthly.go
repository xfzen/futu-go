package akfeeds

import (
	"encoding/json"
	"sort"
	"time"

	"futuq/pkg/gota"
	"futuq/pkg/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type MonthlyData struct {
	// Date  string // 2023-01
	Year  int // 2023
	Month int // 01

	Open   float64 // 开盘价
	Close  float64 // 收盘价
	High   float64 // 最高价
	Low    float64 // 最低价
	Volume float64 // 成交量

	Change float64 // 升跌幅度
}

func HandleIndexData(jsonstr string) error {
	var indexDatas []map[string]interface{}

	err := json.Unmarshal([]byte(jsonstr), &indexDatas)
	if err != nil {
		logx.Errorf("json.Unmarshal err: %v", err)
		return err
	}

	// convert to monthly data
	montylyData := ConvertToMonthlyData(indexDatas)

	dataJsonstr := utils.PrettyJson(montylyData)
	df := gota.DataframeFromJSON(dataJsonstr)

	gota.DemoAkDataFilter(df)

	return nil
}

// 将每日数据转换为每月数据
func ConvertToMonthlyData(dailyData []map[string]interface{}) []MonthlyData {
	logx.Infof("~~~~~~~~~~~~~~~~~~~ConvertToMonthlyData~~~~~~~~~~~~~~~~~~~~~~~~~~")
	logx.Infof("dailyData1 size: %v", len(dailyData))

	// 对每日数据按时间升序排序
	sort.Slice(dailyData, func(i, j int) bool {
		ti, _ := time.Parse("2006-01-02T00:00:00.000Z", dailyData[i]["date"].(string))
		tj, _ := time.Parse("2006-01-02T00:00:00.000Z", dailyData[j]["date"].(string))
		return ti.Before(tj)
	})

	logx.Infof("dailyData2 size: %v", len(dailyData))

	var monthlyData []MonthlyData
	var currentMonth MonthlyData

	for _, data := range dailyData {
		// 解析日期
		dateStr := data["date"].(string)
		date, _ := time.Parse("2006-01-02T00:00:00.000Z", dateStr)

		// 如果当前日期的月份和上一条数据的月份不同，表示进入了新的一个月
		if date.Month() != time.Month(currentMonth.Month) {
			// 如果不是第一条数据，将上一个月的数据添加到结果中
			if currentMonth.Month != 0 {
				monthlyData = append(monthlyData, currentMonth)
			}

			// 计算升跌幅度
			open := data["open"].(float64)
			close := data["close"].(float64)
			change := calcChange(open, close)

			// 初始化新的月份数据
			currentMonth = MonthlyData{
				Year:   date.Year(),
				Month:  int(date.Month()),
				Open:   open,
				High:   data["high"].(float64),
				Low:    data["low"].(float64),
				Close:  close,
				Volume: data["volume"].(float64),
				Change: change,
			}
		} else {
			// 更新当前月份的数据
			currentMonth.High = max(currentMonth.High, data["high"].(float64))
			currentMonth.Low = min(currentMonth.Low, data["low"].(float64))
			open := currentMonth.Open
			close := data["close"].(float64)
			change := calcChange(open, close)

			currentMonth.Close = close
			currentMonth.Volume += data["volume"].(float64)
			currentMonth.Change = change
		}
	}

	// 将最后一个月的数据添加到结果中
	if currentMonth.Month != 0 {
		monthlyData = append(monthlyData, currentMonth)
	}

	return monthlyData
}

// 获取两个数中的最大值
func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

// 获取两个数中的最小值
func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

// 要计算股票的上涨或下跌幅度，您可以使用以下公式：
// 如果涨跌幅为正数，则表示上涨，如果为负数，则表示下跌
func calcChange(open float64, close float64) float64 {
	return (close - open) / open * 100
}
