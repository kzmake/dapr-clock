package service

import (
	"context"

	"github.com/kzmake/dapr-clock/microservices/minute-hand/domain/aggregate"
)

// Movement ...
type Movement interface {
	Increase(context.Context, *aggregate.Hand) (*aggregate.Hand, error)
}
