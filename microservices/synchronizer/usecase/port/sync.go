package port

import (
	"context"
)

// SyncInputData ...
type SyncInputData struct{}

// SyncOutputData ...
type SyncOutputData struct {
	Error error
}

// Sync ...
type Sync interface {
	Handle(context.Context, *SyncInputData) *SyncOutputData
}
