package pubsub

import (
	"context"

	"github.com/dapr/go-sdk/client"
	dapr "github.com/dapr/go-sdk/client"

	// "github.com/kzmake/dapr-clock/microservices/common/domain/event"
	"github.com/kzmake/dapr-clock/microservices/second-hand/domain/aggregate"
	"github.com/kzmake/dapr-clock/microservices/second-hand/domain/service"
	"github.com/kzmake/dapr-clock/microservices/second-hand/domain/vo"
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
	hand := &aggregate.Hand{Second: vo.Second(int(h.Second) + 1)}

	if hand.IsPeriodic() {
		// FIXME: 秒針が1周したタイミングで event.TickTicked60s を pubsub へイベントとして publish する
		return &aggregate.Hand{Second: vo.Second(0)}, nil
	}

	return hand, nil
}
