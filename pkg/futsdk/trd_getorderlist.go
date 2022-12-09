package futuapi

import (
	"context"

	"futuq/pkg/proto/trdcommon"
	"futuq/pkg/proto/trdgetorderlist"
	"futuq/pkg/protocol"
)

const (
	ProtoIDTrdGetOrderList = 2201 // Trd_GetOrderList	获取订单列表
)

// 查询今日订单
func (api *FutuAPI) GetOrderList(ctx context.Context, header *TrdHeader, filter *TrdFilterConditions, status []trdcommon.OrderStatus, refresh bool) ([]*Order, error) {
	req := trdgetorderlist.Request{
		C2S: &trdgetorderlist.C2S{
			Header:           header.pb(),
			FilterConditions: filter.pb(),
			FilterStatusList: orderStatusList(status).pb(),
			RefreshCache:     &refresh,
		},
	}
	ch := make(trdgetorderlist.ResponseChan)
	if err := api.get(ProtoIDTrdGetOrderList, &req, ch); err != nil {
		return nil, err
	}
	select {
	case <-ctx.Done():
		return nil, ErrInterrupted
	case resp, ok := <-ch:
		if !ok {
			return nil, ErrChannelClosed
		}
		return orderListFromPB(resp.GetS2C().GetOrderList()), protocol.Error(resp)
	}
}
