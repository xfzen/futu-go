package futapi

import (
	"context"

	futsdk "futuq/pkg/futsdk"
	"futuq/pkg/gota"
	"futuq/pkg/proto/qotcommon"
	"futuq/pkg/utils"

	log "github.com/pion/ion-log"
)

var (
	opend *futsdk.FutuAPI
	ctx   context.Context
)

// 获取个人分组数据。如可以为‘指数’、‘期货’、‘美股’、‘全部’等
func GetUserSecurity(group string) {
	secs, err := opend.GetUserSecurity(ctx, group)
	if err != nil {
		log.Errorf("GetUserSecurity err: %v", err)
		return
	}

	for _, v := range secs {
		code := v.Basic.Security.Code
		name := v.Basic.Name
		listTime := v.Basic.ListTime
		log.Infof("code: %v, name: %v, listTime: %v", code, name, listTime)
	}
}

// QotMarket_QotMarket_Unknown       QotMarket = 0  //未知市场
// QotMarket_QotMarket_HK_Security   QotMarket = 1  //香港市场
// QotMarket_QotMarket_HK_Future     QotMarket = 2  //港期货(已废弃，使用QotMarket_HK_Security即可)
// QotMarket_QotMarket_US_Security   QotMarket = 11 //美国市场
// QotMarket_QotMarket_CNSH_Security QotMarket = 21 //沪股市场
// QotMarket_QotMarket_CNSZ_Security QotMarket = 22 //深股市场
// QotMarket_QotMarket_SG_Security   QotMarket = 31 //新加坡市场
// QotMarket_QotMarket_JP_Security   QotMarket = 41 //日本市场

// 获取深圳市场股票历史数据
func GetHistData_CNSZ(code string, begin string, end string) {
	seinfo := &futsdk.Security{
		Market: qotcommon.QotMarket_QotMarket_CNSH_Security,
		Code:   code,
	}

	var nextKey []byte

	histData, err := opend.RequestHistoryKLine(ctx, seinfo, begin, end,
		qotcommon.KLType_KLType_Day, qotcommon.RehabType_RehabType_Forward, 5000, qotcommon.KLFields_KLFields_None, nextKey, false)
	if err != nil {
		log.Errorf("RequestHistoryKLine err: %v", err)
		return
	}

	log.Warnf("len(histData): %v, nextKey: %v", len(histData.KLines), nextKey)

	jsonstr := utils.PrettyJson(histData.KLines)

	gota.DemoMtDataFilter(jsonstr)

	// df := dataframe.ReadJSON(strings.NewReader(jsonstr))
	// log.Infof("histData: %v", df)
}
