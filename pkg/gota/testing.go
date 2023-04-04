package gota

import (
	"bytes"

	"futuq/pkg/utils"

	"github.com/go-gota/gota/dataframe"
	"github.com/markcheno/go-talib"
	"github.com/zeromicro/go-zero/core/logx"
)

func DemoMtDataFilter(json string) {
	df := DataframeFromJSON(json)
	logx.Infof(">>DemoMtDataFilter df: %v", df)

	// 收盘数据
	close := df.Col("ClosePrice")

	// 计算均线（120）. 计算均线的时候，数据量需大于均线；不然会崩溃
	ma5 := talib.Sma(close.Float(), 5)

	utils.DumpSlice("\nma5", ma5)

	// test demo
	ma5Delta := calcAng(ma5)

	xlM5 := talib.LinearRegAngle(ma5, 2)
	tanM5 := talib.Tan(ma5Delta)

	utils.DumpSlice("\nxl5", xlM5)
	utils.DumpSlice("tanM5", tanM5)
}

func DemoAkDataFilter(df dataframe.DataFrame) {
	logx.Infof(">>DemoAkDataFilter df: %v", df)

	buf := new(bytes.Buffer)
	_ = df.WriteJSON(buf)

	prettyJson := utils.PrettyJson(buf.String())
	logx.Infof("json: %v", prettyJson)
}
