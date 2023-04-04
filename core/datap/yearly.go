package datap

import (
	"sort"

	"futuq/models"

	"github.com/zeromicro/go-zero/core/logx"
)

// 通过实现sort.Interface接口中的方法来实现排序
type ByYear []YearlyChangeData

func (a ByYear) Len() int           { return len(a) }
func (a ByYear) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByYear) Less(i, j int) bool { return a[i].Year < a[j].Year }

func YearlyDataHandle(dataList []YearlyChangeData) error {
	yearlyDataList := []models.TYearlyData{}

	// 按年份升序排序
	sort.Sort(ByYear(dataList))

	logx.Info("~~~~~~~~~~~~~~~~YearlyDataHandle~~~~~~~~~~~~~~~~~~~~~")
	for _, v := range dataList {
		item := models.TYearlyData{}

		item.Symbol = v.Symbol
		item.Year = v.Year

		item.Jan = v.Jan
		item.Feb = v.Feb
		item.Mar = v.Mar
		item.Apr = v.Apr
		item.May = v.May
		item.Jun = v.Jun
		item.Jul = v.Jul
		item.Aug = v.Aug
		item.Sep = v.Sep
		item.Oct = v.Oct
		item.Nov = v.Nov
		item.Dec = v.Dec

		yearlyDataList = append(yearlyDataList, item)
	}

	// batinsert
	model := models.NewTYealModel()
	err := model.BatchInsert(yearlyDataList)
	if err != nil {
		logx.Error(err)
		return err
	}

	return nil
}

func ProcessMonthlyData(data []MonthlyData) []YearlyChangeData {
	// 使用 map 统计同一 symbol 和年份的月度数据
	yearlyDataMap := make(map[string]map[int]map[int]float64)

	for _, d := range data {
		symbol := d.Symbol
		year := d.Year
		month := d.Month
		change := calcChange(d.Open, d.Close)

		if yearlyDataMap[symbol] == nil {
			yearlyDataMap[symbol] = make(map[int]map[int]float64)
		}

		if yearlyDataMap[symbol][year] == nil {
			yearlyDataMap[symbol][year] = make(map[int]float64)
		}

		yearlyDataMap[symbol][year][month] += change
	}

	// 将统计结果存储到 YearlyChangeData 结构体中
	yearlyData := make([]YearlyChangeData, 0, len(yearlyDataMap))

	for symbol, yearMap := range yearlyDataMap {
		for year, monthMap := range yearMap {
			yearlyDataItem := YearlyChangeData{
				Symbol: symbol,
				Year:   year,
			}

			for month, change := range monthMap {
				switch month {
				case 1:
					yearlyDataItem.Jan = change
				case 2:
					yearlyDataItem.Feb = change
				case 3:
					yearlyDataItem.Mar = change
				case 4:
					yearlyDataItem.Apr = change
				case 5:
					yearlyDataItem.May = change
				case 6:
					yearlyDataItem.Jun = change
				case 7:
					yearlyDataItem.Jul = change
				case 8:
					yearlyDataItem.Aug = change
				case 9:
					yearlyDataItem.Sep = change
				case 10:
					yearlyDataItem.Oct = change
				case 11:
					yearlyDataItem.Nov = change
				case 12:
					yearlyDataItem.Dec = change
				}
			}
			yearlyData = append(yearlyData, yearlyDataItem)
		}
	}

	return yearlyData
}
