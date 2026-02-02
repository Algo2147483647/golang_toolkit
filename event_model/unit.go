package event_model

import (
	"context"
	"github.com/Algo2147483647/golang_toolkit/event_model/eca"
	"github.com/Algo2147483647/golang_toolkit/math/graph"
)

type Unit struct {
	eca.ECA
	Key    string `json:"key"`
	State  string
	Pre    []*Unit                `json:"pre"`
	Post   []*Unit                `json:"post"`
	Params map[string]interface{} `json:"params"`
}

func UnitsToNodes(units []*Unit) []graph.Node {
	result := make([]graph.Node, len(units))
	for i, item := range units {
		units[i] = item
	}
	return result
}

func (unit *Unit) GetPreNodes() []graph.Node {
	return UnitsToNodes(unit.Pre)
}

func (unit *Unit) GetPostNodes() []graph.Node {
	return UnitsToNodes(unit.Post)
}

type UnitInstance struct {
	Key   string `json:"key"`
	State string `json:"state"`
	Unit  *Unit  `json:"unit"`
}

const (
	UnitStateNotStarted    = "not_started"
	UnitStateInProgress    = "in_progress"
	UnitStateConditionPass = "condition_passed"
	UnitStateCompleted     = "completed"
)

func Run(ctx context.Context, units []*Unit, eventTypes []string, req interface{}) error {
	// 1. Get units by events
	newUnits := make([]*Unit, 0)

	for _, unit := range units {
		for _, eventType := range eventTypes {
			if unit.State == UnitStateInProgress && unit.EventTrigger(ctx, eventType) {
				newUnits = append(newUnits, unit)
			}
		}
	}

	units = newUnits

	// 2. Run units
	for _, unit := range units {
		if !unit.ConditionPass(ctx, req) {
			continue
		}

		unit.State = UnitStateConditionPass

		if err := unit.ActionExecute(ctx); err != nil {
			return err
		}

		unit.State = UnitStateCompleted

		for _, item := range unit.Post {
			if item.State == UnitStateNotStarted {
				item.State = UnitStateInProgress
			}

			for _, eventType := range eventTypes {
				if item.State == UnitStateInProgress && item.EventTrigger(ctx, eventType) {
					newUnits = append(newUnits, item)
				}
			}
		}
	}

	return nil
}
