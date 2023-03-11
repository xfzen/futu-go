package akfeeds

import (
	"bytes"
	"encoding/json"

	resty "github.com/go-resty/resty/v2"
	log "github.com/pion/ion-log"
	"github.com/zeromicro/go-zero/core/logx"
)

var _client = resty.New()

func Get(url string, params map[string]string) (r *resty.Response, e error) {
	logx.Infof("methdo=GET, url=%v, params: %v", url, params)

	resp, err := _client.R().
		SetQueryParams(params).
		SetHeader("Accept", "application/json").
		Get(url)

	logx.Infof(">>status: %v", resp.Status())

	return resp, err
}

// post请求封装
func Post(url string, qjson interface{}) (r *resty.Response, e error) {
	// client.SetDebug(true)
	resp, err := _client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(qjson).
		Post(url)

	log.Warnf(">>status: %v", resp.Status())

	return resp, err
}

func _prettyJson(r *resty.Response) {
	body := r.Body()

	var prettyJSON bytes.Buffer
	error := json.Indent(&prettyJSON, body, "", "\t")
	if error != nil {
		logx.Errorf("JSON parse error: %v", error)
		return
	}
}
