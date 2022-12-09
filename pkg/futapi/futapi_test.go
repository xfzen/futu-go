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
	GetHistData_CNSZ("000002", "2021-06-01", "2022-11-17")
}
