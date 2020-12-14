package port

import (
	"context"
)

// SyncInputData ...
type SyncInputData struct {
	Minute int
}

// SyncOutputData ...
type SyncOutputData struct {
	Error error
}

// Sync ...
type Sync interface {
	Handle(context.Context, *SyncInputData) *SyncOutputData
}
