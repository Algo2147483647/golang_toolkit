package event_model

type Event interface {
	GetEventType() string
	GetPayload() interface{}
	IsTriggered() bool
}
