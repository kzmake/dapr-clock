package interactor

import (
	"context"

	"github.com/kzmake/dapr-clock/microservices/minute-hand/domain/repository"
	"github.com/kzmake/dapr-clock/microservices/minute-hand/usecase/port"
)

type now struct {
	handRepository repository.Hand
}

// interfaces
var _ port.Now = (*now)(nil)

// NewNow ...
func NewNow(handRepository repository.Hand) port.Now {
	return &now{handRepository: handRepository}
}

// Handle ...
func (i *now) Handle(ctx context.Context, in *port.NowInputData) *port.NowOutputData {
	current, err := i.handRepository.Get(ctx)
	if err != nil {
		return &port.NowOutputData{Error: err}
	}

	return &port.NowOutputData{CurrentHand: current}
}
