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

func TestGetUserSecurity(t *testing.T) {
	GetUserSecurity("美股")
}

func TestGetHistData_CNSZ(t *testing.T) {
	GetHistData_CNSZ("601238", "2022-09-01", "2023-01-12")
}
