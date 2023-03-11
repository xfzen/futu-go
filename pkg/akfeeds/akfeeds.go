package akfeeds

import (
	"errors"

	"futuq/pkg/gota"

	log "github.com/pion/ion-log"
	"github.com/tidwall/gjson"
	"github.com/zeromicro/go-zero/core/logx"
)

func GetHistIndexDaily(params map[string]string) error {
	baseurl := "http://127.0.0.1:8080/api/public/"

	url := baseurl + "stock_zh_index_daily"
	strResp, err := Get(url, params)
	if err != nil {
		logx.Errorf("%v", err)
		return err
	}

	//  响应报文体, json格式
	jsonstr := strResp.String()

	// dataframe
	df := gota.DataframeFromJSON(jsonstr)
	gota.DemoAkDataFilter(df)

	// handle index data
	HandleIndexData(jsonstr)

	return nil
}

func StockHist(params map[string]string) error {
	baseurl := "http://127.0.0.1:8080/api/public/"

	url := baseurl + "stock_zh_a_hist"
	strResp, err := Get(url, params)
	if err != nil {
		log.Errorf("%v", err)
		return err
	}

	//  响应报文体, json格式
	jsonstr := strResp.String()
	if gjson.Get(jsonstr, "success").String() == "false" {
		log.Errorf("InsertReturnArr failed: %v", err)
		return errors.New(gjson.Get(jsonstr, "errorMessage").String())
	}

	logx.Infof(jsonstr)

	df := gota.DataframeFromJSON(jsonstr)
	gota.DemoAkDataFilter(df)

	return nil
}
