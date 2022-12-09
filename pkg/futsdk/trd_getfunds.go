package futuapi

import (
	"context"

	"futuq/pkg/proto/trdcommon"
	"futuq/pkg/proto/trdgetfunds"
	"futuq/pkg/protocol"
)

const (
	ProtoIDTrdGetFunds = 2101 // Trd_GetFunds	获取账户资金
)

// 查询账户资金
func (api *FutuAPI) GetFunds(ctx context.Context, header *TrdHeader, refresh bool, currency trdcommon.Currency) (*Funds, error) {
	req := trdgetfunds.Request{
		C2S: &trdgetfunds.C2S{
			Header:       header.pb(),
			RefreshCache: &refresh,
		},
	}
	if currency != 0 {
		req.C2S.Currency = (*int32)(&currency)
	}
	ch := make(trdgetfunds.ResponseChan)
	if err := api.get(ProtoIDTrdGetFunds, &req, ch); err != nil {
		return nil, err
	}
	select {
	case <-ctx.Done():
		return nil, ErrInterrupted
	case resp, ok := <-ch:
		if !ok {
			return nil, ErrChannelClosed
		}
		return fundsFromPB(resp.GetS2C().GetFunds()), protocol.Error(resp)
	}
}
