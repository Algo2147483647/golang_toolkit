package event_model

import (
	"context"
	"errors"
)

type ECA struct {
	Events     []Event     `json:"event"`
	Conditions []Condition `json:"condition"`
	Actions    []Action    `json:"action"`
}

func (eca *ECA) Work(ctx context.Context, req interface{}) error {
	pass := false
	for _, event := range eca.Events {
		if event.IsTriggered(ctx, req) {
			pass = true
		}
	}

	if !pass {
		return errors.New("no event triggered")
	}

	for _, condition := range eca.Conditions {
		if !condition.Check() {
			return errors.New("condition check failed")
		}
	}

	for _, action := range eca.Actions {
		err := action.Execute()
		if err != nil {
			return err
		}
	}

	return nil
}
