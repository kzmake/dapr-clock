package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	dapr "github.com/dapr/go-sdk/client"
	"github.com/dapr/go-sdk/service/common"

	"github.com/kzmake/dapr-clock/constants"
)

type synchronizeEvent struct {
	Minute int `json:"minute"`
}

// Synchronize は分針を同期します。
func Synchronize(ctx context.Context, e *common.TopicEvent) (bool, error) {
	log.Printf("subscribe: %s/%s/%s: %s", e.PubsubName, e.Source, e.Topic, e.Data.([]byte))

	client, err := dapr.NewClient()
	if err != nil {
		return false, err
	}

	var payload *synchronizeEvent
	if err := json.Unmarshal(e.Data.([]byte), &payload); err != nil {
		return false, err
	}

	if err := client.SaveState(
		ctx, constants.ComponentStateStore, constants.KeyMinute,
		[]byte(fmt.Sprintf("%d", payload.Minute)),
	); err != nil {
		return false, err
	}

	log.Printf("...%d min", payload.Minute)

	return false, nil
}
