package eca

import "context"

type Event interface {
	GetType() string
	GetAttributes() map[string]interface{}
	IsTrigger(ctx context.Context, req interface{}) bool
}

type EventBase struct {
	Type string
}

func (e *EventBase) GetType() string {
	return e.Type
}

func (e *EventBase) GetAttributes() map[string]interface{} {
	return map[string]interface{}{}
}

func (e *EventBase) IsTrigger(ctx context.Context, req interface{}) bool {
	return false
}
