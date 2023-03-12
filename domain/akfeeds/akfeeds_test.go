package akfeeds

import (
	"flag"
	"testing"

	"futuq/config"
	"futuq/database"

	"github.com/zeromicro/go-zero/core/logx"
)

var testConfigFile = flag.String("ff", "../../etc/futuq-api.yaml", "the config file")

func initDB() {
	logx.Info("~~~~~~~~~~~~~~init~~~~~~~~~~~~~")
	c := config.GetTestConf(*testConfigFile)

	// setup database
	database.SetupTestDatabase(c.Database.Driver, c.Database.Source)
}

func TestGetHistIndexDaily(t *testing.T) {
	initDB()

	params := map[string]string{
		"symbol": "sh000016",
	}

	GetHistIndexDaily(params)
}

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
