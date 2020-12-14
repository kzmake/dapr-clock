package service

import (
	"context"
)

// Movement ...
type Movement interface {
	Tick(context.Context) error
}
