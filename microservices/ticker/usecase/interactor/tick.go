package interactor

import (
	"context"

	"github.com/kzmake/dapr-clock/microservices/ticker/domain/service"
	"github.com/kzmake/dapr-clock/microservices/ticker/usecase/port"
)

type tick struct {
	movementService service.Movement
}

// interfaces
var _ port.Tick = (*tick)(nil)

// NewTick ...
func NewTick(movementService service.Movement) port.Tick {
	return &tick{movementService: movementService}
}

// Handle ...
func (i *tick) Handle(ctx context.Context, in *port.TickInputData) *port.TickOutputData {
	if err := i.movementService.Tick(ctx); err != nil {
		return &port.TickOutputData{Error: err}
	}

	return &port.TickOutputData{}
}
