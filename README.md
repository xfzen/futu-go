# 富途FutuOpenD golang接入

# 简介
* Go语言封装的[富途牛牛OpenAPI](https://openapi.futunn.com/futu-api-doc/)
* 集成go-zero框架，便于对外提供REST api
* 集成gota、dataframe、go-talib对数据进行处理与分析

# 使用说明
`pkg/futapi/futapi_test.go`
```
package futapi

import (
	"context"
	"testing"

	futsdk "futuq/pkg/futsdk"

	log "github.com/pion/ion-log"
)

func init() {
	opend = futsdk.NewFutuAPI()
	ctx = context.TODO()

	err := opend.Connect(ctx, ":11111")
	if err != nil {
		log.Errorf("opend.Connect err: %v", err)
		return
	}
}

// 获取分组数据
func TestGetUserSecurity(t *testing.T) {
	GetUserSecurity("美股")
}

// 获取深市历史数据
func TestGetHistData_CNSZ(t *testing.T) {
	GetHistData_CNSZ("000002", "2021-06-01", "2022-11-17")
}
```

# 版权说明
该项目签署了MIT 授权许可，详情请参阅LICENSE.md

# 鸣谢
* 该项目参考了hurisheng的 [futu-go-api](https://github.com/hurisheng/go-futu-api)
