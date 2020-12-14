package interactor

import (
	"context"

	"github.com/kzmake/dapr-clock/microservices/synchronizer/domain/service"
	"github.com/kzmake/dapr-clock/microservices/synchronizer/usecase/port"
)

type sync struct {
	syncService service.Sync
}

// interfaces
var _ port.Sync = (*sync)(nil)

// NewSync ...
func NewSync(syncService service.Sync) port.Sync {
	return &sync{syncService: syncService}
}

// Handle ...
func (i *sync) Handle(ctx context.Context, in *port.SyncInputData) *port.SyncOutputData {
	if err := i.syncService.Sync(ctx); err != nil {
		return &port.SyncOutputData{Error: err}
	}

	return &port.SyncOutputData{}
}
