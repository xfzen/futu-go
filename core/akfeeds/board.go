package akfeeds

import (
	"futuq/core/datap"
	"futuq/pkg/akhttp"

	"github.com/zeromicro/go-zero/core/logx"
)

func GetSHBoard(params map[string]string) error {
	baseurl := "http://127.0.0.1:8080/api/public/"

	url := baseurl + "stock_sh_a_spot_em"
	strResp, err := akhttp.Get(url, params)
	if err != nil {
		logx.Errorf("%v", err)
		return err
	}

	//  响应报文体, json格式
	jsonstr := strResp.String()

	// logx.Infof(jsonstr)
	datap.HandleBoard("SH", jsonstr)

	return nil
}

func GetSZBoard(params map[string]string) error {
	baseurl := "http://127.0.0.1:8080/api/public/"
	url := baseurl + "stock_sz_a_spot_em"

	strResp, err := akhttp.Get(url, params)
	if err != nil {
		logx.Errorf("%v", err)
		return err
	}

	//  响应报文体, json格式
	jsonstr := strResp.String()

	// logx.Infof(jsonstr)
	datap.HandleBoard("SZ", jsonstr)

	return nil
}
