package akfeeds

// 年变化幅度信息
type YearlyChangeData struct {
	Symbol string
	Year   int

	Jan float64 // 1 月的变化值
	Feb float64 // 2 月的变化值
	Mar float64 // 3 月的变化值
	Apr float64 // 4 月的变化值
	May float64 // 5 月的变化值
	Jun float64 // 6 月的变化值
	Jul float64 // 7 月的变化值
	Aug float64 // 8 月的变化值
	Sep float64 // 9 月的变化值
	Oct float64 // 10 月的变化值
	Nov float64 // 11 月的变化值
	Dec float64 // 12 月的变化值
}

// 月度数据
type MonthlyData struct {
	Symbol string
	Date   string // 2023-01
	Year   int    // 2023
	Month  int    // 01

	Open   float64 // 开盘价
	Close  float64 // 收盘价
	High   float64 // 最高价
	Low    float64 // 最低价
	Volume float64 // 成交量

	Change float64 // 升跌幅度
}
