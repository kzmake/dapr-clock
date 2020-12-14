package pubsub

import (
	"context"

	"github.com/dapr/go-sdk/client"
	dapr "github.com/dapr/go-sdk/client"

	"github.com/kzmake/dapr-clock/microservices/common/domain/event"
	"github.com/kzmake/dapr-clock/microservices/minute-hand/domain/aggregate"
	"github.com/kzmake/dapr-clock/microservices/minute-hand/domain/service"
	"github.com/kzmake/dapr-clock/microservices/minute-hand/domain/vo"
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
func (r *movementService) Increase(ctx context.Context, h *aggregate.Hand) (*aggregate.Hand, error) {
	hand := &aggregate.Hand{Minute: vo.Minute(int(h.Minute) + 1)}

	if hand.IsPeriodic() {
		e := event.Ticked60m{}
		if err := r.client.PublishEvent(ctx, pubsub, event.Topic(e), nil); err != nil {
			return nil, err
		}

		return &aggregate.Hand{Minute: vo.Minute(0)}, nil
	}

	return hand, nil
}
