package port

import (
	"context"

	"github.com/kzmake/dapr-clock/microservices/second-hand/domain/aggregate"
)

// NowInputData ...
type NowInputData struct{}

// NowOutputData ...
type NowOutputData struct {
	CurrentHand *aggregate.Hand
	Error       error
}

// Now ...
type Now interface {
	Handle(context.Context, *NowInputData) *NowOutputData
}
