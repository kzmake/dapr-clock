package aggregate

import (
	"github.com/kzmake/dapr-clock/microservices/minute-hand/domain/vo"
)

const period = 60

// Hand ...
type Hand struct {
	Minute vo.Minute
}

// IsPeriodic ...
func (h *Hand) IsPeriodic() bool { return int(h.Minute)/period == 1 }
