package handler

import (
	"context"
	"encoding/json"
	"log"

	dapr "github.com/dapr/go-sdk/client"
	"github.com/dapr/go-sdk/service/common"
	"github.com/kzmake/dapr-clock/constants"
)

// Now は現在の時刻を取得します。
func Now(ctx context.Context, in *common.InvocationEvent) (*common.Content, error) {
	log.Printf("invocation(now)")

	client, err := dapr.NewClient()
	if err != nil {
		return nil, err
	}

	hRes, err := client.InvokeService(ctx, constants.ServiceHourHand, constants.MethodNow)
	if err != nil {
		return nil, err
	}

	mRes, err := client.InvokeService(ctx, constants.ServiceMinuteHand, constants.MethodNow)
	if err != nil {
		return nil, err
	}

	sRes, err := client.InvokeService(ctx, constants.ServiceSecondHand, constants.MethodNow)
	if err != nil {
		return nil, err
	}

	var h *struct {
		Hour int `json:"hour"`
	}
	if err := json.Unmarshal(hRes, &h); err != nil {
		return nil, err
	}

	var m *struct {
		Minute int `json:"minute"`
	}
	if err := json.Unmarshal(mRes, &m); err != nil {
		return nil, err
	}

	var s *struct {
		Second int `json:"second"`
	}
	if err := json.Unmarshal(sRes, &s); err != nil {
		return nil, err
	}

	payload, err := json.Marshal(&map[string]interface{}{"hour": h.Hour, "minute": m.Minute, "second": s.Second})
	if err != nil {
		return nil, err
	}

	return &common.Content{ContentType: "application/json", Data: payload}, nil
}
