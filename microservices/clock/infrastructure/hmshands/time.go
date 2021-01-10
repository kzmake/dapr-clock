package hmshands

import (
	"context"
	"encoding/json"

	"github.com/dapr/go-sdk/client"
	dapr "github.com/dapr/go-sdk/client"

	"github.com/kzmake/dapr-clock/microservices/clock/domain/aggregate"
	"github.com/kzmake/dapr-clock/microservices/clock/domain/repository"
	"github.com/kzmake/dapr-clock/microservices/clock/domain/vo"
)

type timeRepository struct {
	client client.Client
}

// interfaces
var _ repository.Time = (*timeRepository)(nil)

// NewTimeRepository ...
func NewTimeRepository() (repository.Time, error) {
	client, err := dapr.NewClient()
	if err != nil {
		return nil, err
	}

	return &timeRepository{client: client}, nil
}

// Handle ...
func (r *timeRepository) Get(ctx context.Context) (*aggregate.Time, error) {
	hRes, err := r.client.InvokeService(ctx, "hour-hand", "now")
	if err != nil {
		return nil, err
	}

	mRes, err := r.client.InvokeService(ctx, "minute-hand", "now")
	if err != nil {
		return nil, err
	}

	sRes, err := r.client.InvokeService(ctx, "second-hand", "now")
	if err != nil {
		return nil, err
	}

	var h struct {
		Hour int `json:"hour"`
	}
	if err := json.Unmarshal(hRes, &h); err != nil {
		return nil, err
	}

	var m struct {
		Minute int `json:"minute"`
	}
	if err := json.Unmarshal(mRes, &m); err != nil {
		return nil, err
	}

	var s struct {
		Second int `json:"second"`
	}
	if err := json.Unmarshal(sRes, &s); err != nil {
		return nil, err
	}

	currentTime := &aggregate.Time{
		Hour:   vo.Hour(h.Hour),
		Minute: vo.Minute(m.Minute),
		Second: vo.Second(s.Second),
	}

	return currentTime, nil
}
