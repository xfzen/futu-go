package akfeeds

import "testing"

func TestStockHist(t *testing.T) {
	params := map[string]string{
		"symbol":     "000016",
		"period":     "monthly",
		"start_date": "20000101",
		"end_date":   "20231230",
		"adjust":     "hfq",
	}

	StockHist(params)
}

func TestGetHistIndexDaily(t *testing.T) {
	params := map[string]string{
		"symbol": "sh000016",
	}

	GetHistIndexDaily(params)
}
