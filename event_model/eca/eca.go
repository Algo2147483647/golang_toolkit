package eca

import (
	"context"
	"github.com/Algo2147483647/golang_toolkit/common"
)

type ECA struct {
	Events     []Event     `json:"event"`
	Conditions []Condition `json:"condition"`
	Actions    []Action    `json:"action"`
}

func (eca *ECA) Check(ctx context.Context) bool {
	for _, event := range eca.Events {
		attrsKeys := common.GetKeys(event.GetAttributes(ctx))

		for _, cond := range eca.Conditions {
			for _, attr := range attrsKeys {
				if !common.Contains(cond.GetAttributes(ctx), attr) {
					return false
				}
			}
		}

		for _, action := range eca.Actions {
			for _, attr := range attrsKeys {
				if !common.Contains(action.GetAttributes(ctx), attr) {
					return false
				}
			}
		}
	}

	return true
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
		if !condition.IsPass(ctx) {
			return false
		}
	}

	return true
}

func (eca *ECA) ActionExecute(ctx context.Context) error {
	for _, action := range eca.Actions {
		err := action.Execute(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}
