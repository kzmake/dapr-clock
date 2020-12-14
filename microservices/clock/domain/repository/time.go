package repository

import (
	"context"

	"github.com/kzmake/dapr-clock/microservices/clock/domain/aggregate"
)

// Time ...
type Time interface {
	Get(context.Context) (*aggregate.Time, error)
}
