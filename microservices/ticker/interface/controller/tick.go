package controller

import (
	"context"
	"log"

	"github.com/dapr/go-sdk/service/common"

	"github.com/kzmake/dapr-clock/microservices/ticker/usecase/port"
)

// Tick ...
type Tick interface {
	Tick(context.Context, *common.BindingEvent) ([]byte, error)
}

type tick struct {
	tickInputPort port.Tick
}

// NewTick ...
func NewTick(
	tickInputPort port.Tick,
) Tick {
	return &tick{
		tickInputPort: tickInputPort,
	}
}

// Tick ...
func (c *tick) Tick(ctx context.Context, in *common.BindingEvent) ([]byte, error) {
	log.Printf("binding(tiker): %v", in.Metadata)

	out := c.tickInputPort.Handle(ctx, &port.TickInputData{})
	if err := out.Error; err != nil {
		return nil, err
	}

	return nil, nil
}
