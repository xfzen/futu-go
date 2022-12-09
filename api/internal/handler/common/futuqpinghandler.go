package common

import (
	"net/http"

	"futuq/api/internal/logic/common"
	"futuq/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FutuqPingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := common.NewFutuqPingLogic(r.Context(), svcCtx)
		resp, err := l.FutuqPing()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
