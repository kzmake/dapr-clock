package event

// FIXME: Ticked のイベントを定義する

// Ticked60s ...
type Ticked60s struct{}

// Topic ...
func (t Ticked60s) Topic() string { return "Ticked.60s" }

// Ticked60m ...
type Ticked60m struct{}

// Topic ...
func (t Ticked60m) Topic() string { return "Ticked.60m" }

// Ticked24h ...
type Ticked24h struct{}

// Topic ...
func (t Ticked24h) Topic() string { return "Ticked.24h" }

// Synchronized ...
type Synchronized struct {
	Hour   int `json:"hour"`
	Minute int `json:"minute"`
	Second int `json:"second"`
}

// Topic ...
func (t Synchronized) Topic() string { return "Synchronized" }

// Event ...
type Event interface {
	Topic() string
}

// Topic ...
func Topic(e Event) string { return e.Topic() }
