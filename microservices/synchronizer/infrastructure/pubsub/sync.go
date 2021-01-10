package pubsub

import (
	"context"
	"encoding/json"

	"github.com/beevik/ntp"
	"github.com/dapr/go-sdk/client"
	dapr "github.com/dapr/go-sdk/client"

	"github.com/kzmake/dapr-clock/constants"

	"github.com/kzmake/dapr-clock/microservices/common/domain/event"
	"github.com/kzmake/dapr-clock/microservices/synchronizer/domain/service"
)

const pubsub = "pubsub"

type syncService struct {
	client client.Client
}

// interfaces
var _ service.Sync = (*syncService)(nil)

// NewSyncService ...
func NewSyncService() (service.Sync, error) {
	client, err := dapr.NewClient()
	if err != nil {
		return nil, err
	}

	return &syncService{client: client}, nil
}

// Increase ...
func (r *syncService) Sync(ctx context.Context) error {
	time, err := ntp.Time(constants.NTPServer)
	if err != nil {
		return err
	}
	e := event.Synchronized{
		Hour:   time.Hour(),
		Minute: time.Minute(),
		Second: time.Second(),
	}

	payload, err := json.Marshal(e)
	if err != nil {
		return err
	}

	if err := r.client.PublishEvent(ctx, pubsub, event.Topic(e), payload); err != nil {
		return err
	}

	return nil
}
