package handler

import (
	"context"
	"fmt"
	"log"
	"strconv"

	dapr "github.com/dapr/go-sdk/client"
	"github.com/dapr/go-sdk/service/common"

	"github.com/kzmake/dapr-clock/constants"
)

// Increase は時針を1時間増加させます。
func Increase(ctx context.Context, e *common.TopicEvent) (bool, error) {
	log.Printf("subscribe: %s/%s/%s", e.PubsubName, e.Source, e.Topic)

	client, err := dapr.NewClient()
	if err != nil {
		return false, err
	}

	var hour int
	item, err := client.GetState(ctx, constants.ComponentStateStore, constants.KeyHour)
	if err != nil {
		log.Printf("...initialized hour")
	} else {
		hour, _ = strconv.Atoi(string(item.Value))
		hour, err = increase(ctx, hour)
		if err != nil {
			return false, err
		}
	}

	if err := client.SaveState(
		ctx, constants.ComponentStateStore, constants.KeyHour,
		[]byte(fmt.Sprintf("%d", hour)),
	); err != nil {
		return false, err
	}

	log.Printf("...%d hour", hour)

	return false, nil
}

func increase(ctx context.Context, n int) (int, error) {
	if (n+1)/24 == 1 {
		client, err := dapr.NewClient()
		if err != nil {
			return 0, err
		}
		if err := client.PublishEvent(ctx, constants.ComponentPubSub, constants.EventTicked24h, nil); err != nil {
			return 0, err
		}
	}

	log.Printf("...published %s", constants.EventTicked24h)

	return (n + 1) % 24, nil
}
