package port

import (
	"context"
)

// IncreaseInputData ...
type IncreaseInputData struct{}

// IncreaseOutputData ...
type IncreaseOutputData struct {
	Error error
}

// Increase ...
type Increase interface {
	Handle(context.Context, *IncreaseInputData) *IncreaseOutputData
}
