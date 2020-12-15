package repository

import (
	"context"

	"github.com/kzmake/dapr-clock/microservices/hour-hand/domain/aggregate"
)

// Hand ...
type Hand interface {
	Get(context.Context) (*aggregate.Hand, error)
	Set(context.Context, *aggregate.Hand) error
}
