package statestore

import (
	"context"
	"fmt"
	"strconv"

	"github.com/dapr/go-sdk/client"
	dapr "github.com/dapr/go-sdk/client"

	"github.com/kzmake/dapr-clock/microservices/hour-hand/domain/aggregate"
	"github.com/kzmake/dapr-clock/microservices/hour-hand/domain/repository"
	"github.com/kzmake/dapr-clock/microservices/hour-hand/domain/vo"
)

const (
	statestore = "statestore"
	key        = "hour"
)

type handRepository struct {
	client client.Client
}

// interfaces
var _ repository.Hand = (*handRepository)(nil)

// NewHandRepository ...
func NewHandRepository() (repository.Hand, error) {
	client, err := dapr.NewClient()
	if err != nil {
		return nil, err
	}

	return &handRepository{client: client}, nil
}

// Get ...
func (r *handRepository) Get(ctx context.Context) (*aggregate.Hand, error) {
	item, err := r.client.GetState(ctx, statestore, key)
	if err != nil {
		return nil, err
	}

	hour, err := strconv.Atoi(string(item.Value))
	if err != nil {
		return nil, err
	}

	return &aggregate.Hand{Hour: vo.Hour(hour)}, nil
}

// Set ...
func (r *handRepository) Set(ctx context.Context, hand *aggregate.Hand) error {
	if err := r.client.SaveState(ctx, statestore, key, []byte(fmt.Sprintf("%d", hand.Hour))); err != nil {
		return err
	}

	return nil
}
