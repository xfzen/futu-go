package futuapi

import (
	"context"

	"futuq/pkg/proto/trdcommon"
	"futuq/pkg/proto/trdgethistoryorderlist"
	"futuq/pkg/protocol"
)

const (
	ProtoIDTrdGetHistoryOrderList = 2221 // Trd_GetHistoryOrderList	获取历史订单列表
)

// 查询历史订单
func (api *FutuAPI) GetHistoryOrderList(ctx context.Context, header *TrdHeader, filter *TrdFilterConditions, status []trdcommon.OrderStatus) ([]*Order, error) {
	req := trdgethistoryorderlist.Request{
		C2S: &trdgethistoryorderlist.C2S{
			Header:           header.pb(),
			FilterConditions: filter.pb(),
			FilterStatusList: orderStatusList(status).pb(),
		},
	}
	ch := make(trdgethistoryorderlist.ResponseChan)
	if err := api.get(ProtoIDTrdGetHistoryOrderList, &req, ch); err != nil {
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
