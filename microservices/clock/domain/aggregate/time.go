package aggregate

import (
	"github.com/kzmake/dapr-clock/microservices/clock/domain/vo"
)

// Time ...
type Time struct {
	Hour   vo.Hour
	Minute vo.Minute
	Second vo.Second
}
