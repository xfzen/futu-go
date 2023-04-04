package datap

import (
	"encoding/json"

	"futuq/models"

	"github.com/zeromicro/go-zero/core/logx"
)

// 处理沪市板块数据（用于收盘后获取个股实时数据）
func HandleBoard(mkt string, jsonstr string) error {
	var indexSpot []map[string]interface{}

	boardDataList := []models.TBoardInfoData{}

	err := json.Unmarshal([]byte(jsonstr), &indexSpot)
	if err != nil {
		logx.Errorf("json.Unmarshal err: %v", err)
		return err
	}

	// map[5分钟涨跌:-0.63 60日涨跌幅:-7.1 今开:63.11 代码:688212 名称:澳华内镜 市净率:6.69 市盈率-动态:413.29 年初至今涨跌幅:-2.89
	// 序号:448 总市值:8.480424e+09 成交量:20721 成交额:1.30009598e+08 振幅:4.12
	// 换手率:2.31 昨收:63.1 最低:61.4 最新价:63.6 最高:64 流通市值:5.71447272e+09 涨跌幅:0.79 涨跌额:0.5 涨速:-0.31 量比:1.44]
	for _, v := range indexSpot {
		symbol := v["代码"]
		name := v["名称"]
		close := v["昨收"]
		high := v["最高"]
		low := v["最低"]

		// 对异常数据进行过滤
		if close == nil || high == nil || low == nil {
			continue
		}

		item := models.TBoardInfoData{}
		item.Symbol = symbol.(string)
		item.Name = name.(string)
		item.Close = close.(float64)
		item.High = high.(float64)
		item.Low = low.(float64)
		item.Market = mkt

		boardDataList = append(boardDataList, item)
	}

	// batinsert
	model := models.NewTBoardInfoModel()
	err = model.BatchInsert(boardDataList)
	if err != nil {
		logx.Error(err)
		return err
	}

	return nil
}
