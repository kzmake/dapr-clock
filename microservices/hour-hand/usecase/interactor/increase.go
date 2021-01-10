package interactor

import (
	"context"

	"github.com/kzmake/dapr-clock/microservices/hour-hand/domain/repository"
	"github.com/kzmake/dapr-clock/microservices/hour-hand/domain/service"
	"github.com/kzmake/dapr-clock/microservices/hour-hand/usecase/port"
)

type increase struct {
	handRepository  repository.Hand
	movementService service.Movement
}

// interfaces
var _ port.Increase = (*increase)(nil)

// NewIncrease ...
func NewIncrease(
	handRepository repository.Hand,
	movementService service.Movement,
) port.Increase {
	return &increase{
		handRepository:  handRepository,
		movementService: movementService,
	}
}

// Handle ...
func (i *increase) Handle(ctx context.Context, in *port.IncreaseInputData) *port.IncreaseOutputData {
	_current, err := i.handRepository.Get(ctx)
	if err != nil {
		return &port.IncreaseOutputData{Error: err}
	}

	current, err := i.movementService.Increase(ctx, _current)
	if err != nil {
		return &port.IncreaseOutputData{Error: err}
	}

	if err := i.handRepository.Set(ctx, current); err != nil {
		return &port.IncreaseOutputData{Error: err}
	}

	return &port.IncreaseOutputData{}
}
