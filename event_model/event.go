package event_model

type EventEntity struct {
	Events     []Event     `json:"event"`
	Conditions []Condition `json:"condition"`
	Actions    []Action    `json:"action"`
}

type Event interface {
	GetEventType() string
	GetPayload() interface{}
}

type Condition interface {
	Check() bool
}

type Action interface {
	Execute() error
}
