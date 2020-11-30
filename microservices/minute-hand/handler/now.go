package handler

import (
	"context"
	"encoding/json"
	"log"
	"strconv"

	dapr "github.com/dapr/go-sdk/client"
	"github.com/dapr/go-sdk/service/common"
	"github.com/kzmake/dapr-clock/constants"
)

type nowResponse struct {
	Minute int `json:"minute"`
}

// Now は現在の秒針を取得します。
func Now(ctx context.Context, in *common.InvocationEvent) (*common.Content, error) {
	log.Printf("invocation(now)")

	client, err := dapr.NewClient()
	if err != nil {
		return nil, err
	}

	item, err := client.GetState(ctx, constants.ComponentStateStore, constants.KeyMinute)
	if err != nil {
		return nil, err
	}

	minute, err := strconv.Atoi(string(item.Value))
	if err != nil {
		return nil, err
	}

	payload, err := json.Marshal(&nowResponse{Minute: minute})
	if err != nil {
		return nil, err
	}

	return &common.Content{ContentType: "application/json", Data: payload}, nil
}