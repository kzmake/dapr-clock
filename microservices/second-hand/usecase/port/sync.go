package port

import (
	"context"
)

// SyncInputData ...
type SyncInputData struct {
	Second int
}

// SyncOutputData ...
type SyncOutputData struct {
	Error error
}

// Sync ...
type Sync interface {
	Handle(context.Context, *SyncInputData) *SyncOutputData
}
