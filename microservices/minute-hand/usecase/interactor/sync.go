package interactor

import (
	"context"

	"github.com/kzmake/dapr-clock/microservices/minute-hand/domain/aggregate"
	"github.com/kzmake/dapr-clock/microservices/minute-hand/domain/repository"
	"github.com/kzmake/dapr-clock/microservices/minute-hand/domain/vo"
	"github.com/kzmake/dapr-clock/microservices/minute-hand/usecase/port"
)

type sync struct {
	handRepository repository.Hand
}

// interfaces
var _ port.Sync = (*sync)(nil)

// NewSync ...
func NewSync(handRepository repository.Hand) port.Sync {
	return &sync{handRepository: handRepository}
}

// Handle ...
func (i *sync) Handle(ctx context.Context, in *port.SyncInputData) *port.SyncOutputData {
	current := &aggregate.Hand{Minute: vo.Minute(in.Minute)}

	if err := i.handRepository.Set(ctx, current); err != nil {
		return &port.SyncOutputData{Error: err}
	}

	return &port.SyncOutputData{}
}
