package handler

import (
	"context"
	"encoding/json"
	"log"

	"github.com/beevik/ntp"
	dapr "github.com/dapr/go-sdk/client"
	"github.com/dapr/go-sdk/service/common"

	"github.com/kzmake/dapr-clock/constants"
)

type synchronizeEvent struct {
	Hour   int `json:"hour"`
	Minute int `json:"minute"`
	Second int `json:"second"`
}

// Synchronize は EventSynchronized を発行します。
func Synchronize(ctx context.Context, in *common.BindingEvent) ([]byte, error) {
	log.Printf("binding(synchronizer): %v", in.Metadata)

	client, err := dapr.NewClient()
	if err != nil {
		return nil, err
	}

	time, err := ntp.Time(constants.NTPServer)
	if err != nil {
		return nil, err
	}

	payload, err := json.Marshal(&synchronizeEvent{Hour: time.Hour(), Minute: time.Minute(), Second: time.Second()})
	if err != nil {
		return nil, err
	}

	if err := client.PublishEvent(ctx, constants.ComponentPubSub, constants.EventSynchronized, payload); err != nil {
		return nil, err
	}

	log.Printf("...published %s: %s", constants.EventSynchronized, string(payload))

	return nil, nil
}
