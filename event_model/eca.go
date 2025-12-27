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
	pass := false
	for _, event := range eca.Events {
		if event.IsTrigger(ctx, req) {
			pass = true
		}
	}
	return pass
}

func (eca *ECA) ConditionPass(ctx context.Context, req interface{}) (bool, error) {
	for _, condition := range eca.Conditions {
		pass, err := condition.Check()
		if err != nil {
			return false, err
		} else if !pass {
			return false, nil
		}
	}

	return true, nil
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
