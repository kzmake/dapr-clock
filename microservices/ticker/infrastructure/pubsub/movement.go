package pubsub

import (
	"context"

	"github.com/dapr/go-sdk/client"
	dapr "github.com/dapr/go-sdk/client"

	// "github.com/kzmake/dapr-clock/microservices/common/domain/event"
	"github.com/kzmake/dapr-clock/microservices/ticker/domain/service"
)

const pubsub = "pubsub"

type movementService struct {
	client client.Client
}

// interfaces
var _ service.Movement = (*movementService)(nil)

// NewMovementService ...
func NewMovementService() (service.Movement, error) {
	client, err := dapr.NewClient()
	if err != nil {
		return nil, err
	}

	return &movementService{client: client}, nil
}

// Increase ...
func (r *movementService) Tick(ctx context.Context) error {
	// FIXME: event.Ticked を定義して pubsub へイベントとして publish する
	return nil
}
