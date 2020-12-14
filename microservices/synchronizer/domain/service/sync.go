package service

import (
	"context"
)

// Sync ...
type Sync interface {
	Sync(context.Context) error
}
