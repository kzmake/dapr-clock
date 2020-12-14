package controller

import (
	"context"
	"log"

	"github.com/dapr/go-sdk/service/common"
	"github.com/kzmake/dapr-clock/microservices/synchronizer/usecase/port"
)

// Sync ...
type Sync interface {
	Sync(context.Context, *common.BindingEvent) ([]byte, error)
}

type sync struct {
	syncInputPort port.Sync
}

// NewSync ...
func NewSync(
	syncInputPort port.Sync,
) Sync {
	return &sync{
		syncInputPort: syncInputPort,
	}
}

// Sync ...
func (c *sync) Sync(ctx context.Context, in *common.BindingEvent) ([]byte, error) {
	log.Printf("binding(synchronizer): %v", in.Metadata)

	out := c.syncInputPort.Handle(ctx, &port.SyncInputData{})
	if err := out.Error; err != nil {
		return nil, err
	}

	return nil, nil
}
