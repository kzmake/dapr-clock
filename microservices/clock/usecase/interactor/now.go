package interactor

import (
	"context"
	"log"

	"github.com/kzmake/dapr-clock/microservices/clock/domain/repository"
	"github.com/kzmake/dapr-clock/microservices/clock/usecase/port"
)

type now struct {
	timeRepository repository.Time
}

// interfaces
var _ port.Now = (*now)(nil)

// NewNow ...
func NewNow(timeRepository repository.Time) port.Now {
	return &now{timeRepository: timeRepository}
}

// Handle ...
func (i *now) Handle(ctx context.Context, in *port.NowInputData) *port.NowOutputData {
	log.Printf("invocation(now)")

	currentTime, err := i.timeRepository.Get(ctx)
	if err != nil {
		return &port.NowOutputData{Error: err}
	}

	return &port.NowOutputData{CurrentTime: currentTime}
}
