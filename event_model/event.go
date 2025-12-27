package event_model

type Event interface {
	GetEventType() string
	GetPayload() interface{}
	IsTriggered() bool
}

type EventBase struct {
}

func (e *EventBase) GetEventType() string {
	return "EventBase"
}

func (e *EventBase) GetPayload() interface{} {
	return nil
}

func (e *EventBase) IsTriggered() bool {
	return false
}
