package event_model

import "context"

type Event interface {
	GetEventType() string
	GetPayload() interface{}
	IsTriggered(ctx context.Context, req interface{}) bool
}

type EventBase struct {
}

func (e *EventBase) GetEventType() string {
	return "EventBase"
}

func (e *EventBase) GetPayload() interface{} {
	return nil
}

func (e *EventBase) IsTriggered(ctx context.Context, req interface{}) bool {
	return false
}
