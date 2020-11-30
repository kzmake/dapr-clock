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

// Increase は秒針を1秒間増加させます。
func Increase(ctx context.Context, e *common.TopicEvent) (bool, error) {
	log.Printf("subscribe: %s/%s/%s", e.PubsubName, e.Source, e.Topic)

	client, err := dapr.NewClient()
	if err != nil {
		return false, err
	}

	var sec int
	item, err := client.GetState(ctx, constants.ComponentStateStore, constants.KeySecond)
	if err != nil {
		log.Printf("...initialized")
	} else {
		sec, _ = strconv.Atoi(string(item.Value))
		sec, err = increase(ctx, sec)
		if err != nil {
			return false, err
		}
	}

	if err := client.SaveState(
		ctx, constants.ComponentStateStore, constants.KeySecond,
		[]byte(fmt.Sprintf("%d", sec)),
	); err != nil {
		return false, err
	}

	log.Printf("...%d sec", sec)

	return false, nil
}

func increase(ctx context.Context, n int) (int, error) {
	if (n+1)/60 == 1 {
		client, err := dapr.NewClient()
		if err != nil {
			return 0, err
		}
		if err := client.PublishEvent(ctx, constants.ComponentPubSub, constants.EventTicked60s, nil); err != nil {
			return 0, err
		}

		log.Printf("...published %s", constants.EventTicked60s)
	}

	return (n + 1) % 60, nil
}
