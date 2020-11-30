package handler

import (
	"context"
	"log"

	dapr "github.com/dapr/go-sdk/client"
	"github.com/dapr/go-sdk/service/common"

	"github.com/kzmake/dapr-clock/constants"
)

// Tick は EventTicked を発行します。
func Tick(ctx context.Context, in *common.BindingEvent) ([]byte, error) {
	log.Printf("binding(tiker): %v", in.Metadata)

	client, err := dapr.NewClient()
	if err != nil {
		return nil, err
	}

	if err := client.PublishEvent(ctx, constants.ComponentPubSub, constants.EventTicked, nil); err != nil {
		return nil, err
	}

	log.Printf("...published %s", constants.EventTicked)

	return nil, nil
}
