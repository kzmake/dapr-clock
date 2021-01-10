package pubsub

import (
	"context"

	"github.com/dapr/go-sdk/client"
	dapr "github.com/dapr/go-sdk/client"

	"github.com/kzmake/dapr-clock/microservices/common/domain/event"

	"github.com/kzmake/dapr-clock/microservices/hour-hand/domain/aggregate"
	"github.com/kzmake/dapr-clock/microservices/hour-hand/domain/service"
	"github.com/kzmake/dapr-clock/microservices/hour-hand/domain/vo"
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
	hand := &aggregate.Hand{Hour: vo.Hour(int(h.Hour) + 1)}

	if hand.IsPeriodic() {
		e := event.Ticked24h{}
		if err := r.client.PublishEvent(ctx, pubsub, event.Topic(e), nil); err != nil {
			return nil, err
		}

		return &aggregate.Hand{Hour: vo.Hour(0)}, nil
	}

	return hand, nil
}
