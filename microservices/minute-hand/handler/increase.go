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

// Increase は分針を1分間増加させます。
func Increase(ctx context.Context, e *common.TopicEvent) (bool, error) {
	log.Printf("subscribe: %s/%s/%s", e.PubsubName, e.Source, e.Topic)

	client, err := dapr.NewClient()
	if err != nil {
		return false, err
	}

	var min int
	item, err := client.GetState(ctx, constants.ComponentStateStore, constants.KeyMinute)
	if err != nil {
		log.Printf("...initialized min")
	} else {
		min, _ = strconv.Atoi(string(item.Value))
		min, err = increase(ctx, min)
		if err != nil {
			return false, err
		}
	}

	if err := client.SaveState(
		ctx, constants.ComponentStateStore, constants.KeyMinute,
		[]byte(fmt.Sprintf("%d", min)),
	); err != nil {
		return false, err
	}

	log.Printf("...%d min", min)

	return false, nil
}

func increase(ctx context.Context, n int) (int, error) {
	if (n+1)/60 == 1 {
		client, err := dapr.NewClient()
		if err != nil {
			return 0, err
		}
		if err := client.PublishEvent(ctx, constants.ComponentPubSub, constants.EventTicked60m, nil); err != nil {
			return 0, err
		}

		log.Printf("...published %s", constants.EventTicked60m)
	}

	return (n + 1) % 60, nil
}
