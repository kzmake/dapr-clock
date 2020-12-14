package statestore

import (
	"context"
	"fmt"
	"strconv"

	"github.com/dapr/go-sdk/client"
	dapr "github.com/dapr/go-sdk/client"

	"github.com/kzmake/dapr-clock/microservices/second-hand/domain/aggregate"
	"github.com/kzmake/dapr-clock/microservices/second-hand/domain/repository"
	"github.com/kzmake/dapr-clock/microservices/second-hand/domain/vo"
)

const (
	statestore = "statestore"
	key        = "second"
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

	second, err := strconv.Atoi(string(item.Value))
	if err != nil {
		return nil, err
	}

	return &aggregate.Hand{Second: vo.Second(second)}, nil
}

// Set ...
func (r *handRepository) Set(ctx context.Context, hand *aggregate.Hand) error {
	if err := r.client.SaveState(ctx, statestore, key, []byte(fmt.Sprintf("%d", hand.Second))); err != nil {
		return err
	}

	return nil
}
