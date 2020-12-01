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

	payload, err := json.Marshal(map[string]interface{}{"hour": time.Hour(), "minute": time.Minute(), "second": time.Second()})
	if err != nil {
		return nil, err
	}

	if err := client.PublishEvent(ctx, constants.ComponentPubSub, constants.EventSynchronized, payload); err != nil {
		return nil, err
	}

	log.Printf("...published %s: %s", constants.EventSynchronized, string(payload))

	return nil, nil
}
