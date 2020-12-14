package controller

import (
	"context"
	"encoding/json"
	"log"

	"github.com/dapr/go-sdk/service/common"
	"github.com/kzmake/dapr-clock/microservices/common/domain/event"
	"github.com/kzmake/dapr-clock/microservices/second-hand/usecase/port"
)

// Hand ...
type Hand interface {
	Now(context.Context, *common.InvocationEvent) (*common.Content, error)
	Increase(context.Context, *common.TopicEvent) (bool, error)
	Sync(context.Context, *common.TopicEvent) (bool, error)
}

type hand struct {
	nowInputPort      port.Now
	increaseInputPort port.Increase
	syncInputPort     port.Sync
}

// NewHand ...
func NewHand(
	nowInputPort port.Now,
	increaseInputPort port.Increase,
	syncInputPort port.Sync,
) Hand {
	return &hand{
		nowInputPort:      nowInputPort,
		increaseInputPort: increaseInputPort,
		syncInputPort:     syncInputPort,
	}
}

// Now ...
func (c *hand) Now(ctx context.Context, in *common.InvocationEvent) (*common.Content, error) {
	log.Printf("invocation(now)")

	out := c.nowInputPort.Handle(ctx, &port.NowInputData{})
	if err := out.Error; err != nil {
		return nil, err
	}

	payload, err := json.Marshal(&map[string]interface{}{
		"second": out.CurrentHand.Second,
	})
	if err != nil {
		return nil, err
	}

	return &common.Content{ContentType: "application/json", Data: payload}, nil
}

// Increase ...
func (c *hand) Increase(ctx context.Context, e *common.TopicEvent) (bool, error) {
	log.Printf("subscribe: %s/%s/%s", e.PubsubName, e.Source, e.Topic)

	out := c.increaseInputPort.Handle(ctx, &port.IncreaseInputData{})
	if err := out.Error; err != nil {
		return false, err
	}

	return true, nil
}

// Sync ...
func (c *hand) Sync(ctx context.Context, e *common.TopicEvent) (bool, error) {
	log.Printf("subscribe: %s/%s/%s", e.PubsubName, e.Source, e.Topic)

	var s event.Synchronized
	if err := json.Unmarshal(e.Data.([]byte), &s); err != nil {
		return false, err
	}

	out := c.syncInputPort.Handle(ctx, &port.SyncInputData{Second: s.Second})
	if err := out.Error; err != nil {
		return false, err
	}

	return true, nil
}
