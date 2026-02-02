package eca

import "context"

type Event interface {
	GetType(ctx context.Context) string
	GetAttributes(ctx context.Context) map[string]interface{}
	IsTrigger(ctx context.Context, req interface{}) bool
}

type EventBase struct {
	Type       string
	Attributes map[string]interface{}
}

func (e *EventBase) GetType(ctx context.Context) string {
	return e.Type
}

func (e *EventBase) GetAttributes(ctx context.Context) map[string]interface{} {
	return e.Attributes
}

func (e *EventBase) IsTrigger(ctx context.Context, req interface{}) bool {
	return false
}
