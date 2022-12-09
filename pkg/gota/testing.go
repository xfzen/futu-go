package gota

import (
	"futuq/pkg/utils"

	"github.com/go-gota/gota/dataframe"
	"github.com/markcheno/go-talib"
	log "github.com/pion/ion-log"
)

func DemoMtDataFilter(json string) {
	df := DataframeFromJSON(json)
	log.Infof(">>DemoMtDataFilter df: %v", df)

	// log.Warnf("json: %v", json)

	// 收盘数据
	close := df.Col("ClosePrice")

	// 计算均线（120）. 计算均线的时候，数据量需大于均线；不然会崩溃
	ma20 := talib.Sma(close.Float(), 20)
	ma120 := talib.Sma(close.Float(), 120)
	ma250 := talib.Sma(close.Float(), 250)

	utils.DumpSlice("\nma020", ma20)
	utils.DumpSlice("ma120", ma120)
	utils.DumpSlice("ma250", ma250)

	xlM20 := talib.LinearRegAngle(ma20, 10)
	xlM120 := talib.LinearRegAngle(ma120, 10)
	xlM250 := talib.LinearRegAngle(ma250, 10)

	utils.DumpSlice("\nxl020", xlM20)
	utils.DumpSlice("xl120", xlM120)
	utils.DumpSlice("xl250", xlM250)
}

func DemoAkDataFilter(df dataframe.DataFrame) {
	log.Infof(">>DemoAkDataFilter df: %v", df)

	// fil := df.Filter(
	// 	dataframe.F{0, "开盘", series.Greater, "26.530000"},
	// 	// dataframe.F{"A", series.Eq, "a"},
	// 	// dataframe.F{"B", series.Greater, 4},
	// )

	// 收盘数据
	close := df.Col("ClosePrice")
	log.Warnf(">>close: %v", close)

	// 计算均线（5）
	// 计算均线的时候，数据量需大于均线；不然会崩溃
	sma5 := talib.Sma(close.Float(), 5)
	log.Infof(">>colName: %v", df.Names())
	log.Infof(">>sma5: %v", sma5)
}
