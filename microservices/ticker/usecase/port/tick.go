package port

import (
	"context"
)

// TickInputData ...
type TickInputData struct {
	Second int
}

// TickOutputData ...
type TickOutputData struct {
	Error error
}

// Tick ...
type Tick interface {
	Handle(context.Context, *TickInputData) *TickOutputData
}
