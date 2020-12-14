package port

import (
	"context"

	"github.com/kzmake/dapr-clock/microservices/clock/domain/aggregate"
)

// NowInputData ...
type NowInputData struct{}

// NowOutputData ...
type NowOutputData struct {
	CurrentTime *aggregate.Time
	Error       error
}

// Now ...
type Now interface {
	Handle(context.Context, *NowInputData) *NowOutputData
}
