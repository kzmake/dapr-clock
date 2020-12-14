package aggregate

import (
	"github.com/kzmake/dapr-clock/microservices/hour-hand/domain/vo"
)

const period = 24

// Hand ...
type Hand struct {
	Hour vo.Hour
}

// IsPeriodic ...
func (h *Hand) IsPeriodic() bool { return int(h.Hour)/period == 1 }
