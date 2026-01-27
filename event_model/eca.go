package event_model

import (
	"context"
)

type ECA struct {
	Events     []Event     `json:"event"`
	Conditions []Condition `json:"condition"`
	Actions    []Action    `json:"action"`
}

func (eca *ECA) EventTrigger(ctx context.Context, req interface{}) bool {
	for _, event := range eca.Events {
		if event.IsTrigger(ctx, req) {
			return true
		}
	}
	return false
}

func (eca *ECA) ConditionPass(ctx context.Context, req interface{}) bool {
	for _, condition := range eca.Conditions {
		if !condition.IsPass() {
			return false
		}
	}

	return true
}

func (eca *ECA) ActionExecute() error {
	for _, action := range eca.Actions {
		err := action.Execute()
		if err != nil {
			return err
		}
	}
	return nil
}
