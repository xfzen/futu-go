package akfeeds

import (
	"errors"

	"futuq/core/datap"
	"futuq/pkg/akhttp"
	"futuq/pkg/gota"

	"github.com/tidwall/gjson"
	"github.com/zeromicro/go-zero/core/logx"
)

// 指数实时数据
func GetSpotIndex(params map[string]string) error {
	baseurl := "http://127.0.0.1:8080/api/public/"

	url := baseurl + "stock_zh_index_spot"
	strResp, err := akhttp.Get(url, params)
	if err != nil {
		logx.Errorf("%v", err)
		return err
	}

	//  响应报文体, json格式
	jsonstr := strResp.String()

	// logx.Infof(jsonstr)
	datap.HandleIndexSpot(jsonstr)

	return nil
}

// 历史指数行情
func GetHistIndexDaily(params map[string]string) error {
	baseurl := "http://127.0.0.1:8080/api/public/"

	url := baseurl + "stock_zh_index_daily"
	strResp, err := akhttp.Get(url, params)
	if err != nil {
		logx.Errorf("%v", err)
		return err
	}

	//  响应报文体, json格式
	jsonstr := strResp.String()
	symbol := params["symbol"]

	// dataframe
	// df := gota.DataframeFromJSON(jsonstr)
	// gota.DemoAkDataFilter(df)

	// handle index data
	datap.HandleIndexData(symbol, jsonstr)

	return nil
}

func StockHist(params map[string]string) error {
	baseurl := "http://127.0.0.1:8080/api/public/"

	url := baseurl + "stock_zh_a_hist"
	strResp, err := akhttp.Get(url, params)
	if err != nil {
		logx.Errorf("%v", err)
		return err
	}

	//  响应报文体, json格式
	jsonstr := strResp.String()
	if gjson.Get(jsonstr, "success").String() == "false" {
		logx.Errorf("InsertReturnArr failed: %v", err)
		return errors.New(gjson.Get(jsonstr, "errorMessage").String())
	}

	logx.Infof(jsonstr)

	df := gota.DataframeFromJSON(jsonstr)
	gota.DemoAkDataFilter(df)

	return nil
}
