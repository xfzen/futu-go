package futuapi

import (
	"context"
	"testing"

	"github.com/zeromicro/go-zero/core/logx"
)

func TestConnect(t *testing.T) {
	api := NewFutuAPI()
	defer api.Close(context.Background())

	// api.SetRecvNotify(true)
	// nCh, err := api.SysNotify(context.Background())
	// if err != nil {
	// 	t.Error(err)
	// }

	err := api.Connect(context.Background(), ":11111")
	if err != nil {
		logx.Errorf("api.Connect err: %v", err)
		return
	}

	logx.Infof(">> err: %v", err)

	// if sub, err := api.QuerySubscription(context.Background(), true); err != nil {
	// 	t.Error(err)
	// } else {
	// 	t.Error(sub)
	// }

	// tCh, err := api.UpdateTicker(context.Background())
	// if err != nil {
	// 	t.Error(err)
	// }
	// if err := api.Subscribe(context.Background(), []*Security{{qotcommon.QotMarket_QotMarket_HK_Security, "00700"}}, []qotcommon.SubType{qotcommon.SubType_SubType_Ticker}, true, true, true, true); err != nil {
	// 	t.Error(err)
	// }

	// select {
	// case notify := <-nCh:
	// 	log.Warnf(">>notify: %v", notify)

	// case ticker := <-tCh:
	// 	log.Warnf(">>ticker: %v", ticker)
	// }

	// if sub, err := api.QuerySubscription(context.Background(), true); err != nil {
	// 	t.Error(err)
	// } else {
	// 	t.Error(sub)
	// }

	// secs, err := api.GetUserSecurity(context.Background(), "全部")
	// if err != nil {
	// 	t.Error(err)
	// } else {
	// 	t.Error(secs)
	// }
}
