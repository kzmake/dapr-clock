package aggregate

import (
	"github.com/kzmake/dapr-clock/microservices/second-hand/domain/vo"
)

const period = 60

// Hand ...
type Hand struct {
	Second vo.Second
}

// IsPeriodic ...
func (h *Hand) IsPeriodic() bool { return int(h.Second)/period == 1 }
