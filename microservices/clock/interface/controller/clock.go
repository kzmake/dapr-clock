package controller

import (
	"context"
	"encoding/json"

	"github.com/dapr/go-sdk/service/common"
	"github.com/kzmake/dapr-clock/microservices/clock/usecase/port"
)

// Clock ...
type Clock interface {
	Now(context.Context, *common.InvocationEvent) (*common.Content, error)
}

type clock struct {
	nowInputPort port.Now
}

// NewClock ...
func NewClock(nowInputPort port.Now) Clock {
	return &clock{nowInputPort: nowInputPort}
}

// Now ...
func (c *clock) Now(ctx context.Context, in *common.InvocationEvent) (*common.Content, error) {
	out := c.nowInputPort.Handle(ctx, &port.NowInputData{})
	if err := out.Error; err != nil {
		return nil, err
	}

	payload, err := json.Marshal(&map[string]interface{}{
		"hour":   out.CurrentTime.Hour,
		"minute": out.CurrentTime.Minute,
		"second": out.CurrentTime.Second,
	})
	if err != nil {
		return nil, err
	}

	return &common.Content{ContentType: "application/json", Data: payload}, nil
}
