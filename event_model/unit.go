package event_model

import (
	"github.com/Algo2147483647/golang_toolkit/event_model/eca"
	"github.com/Algo2147483647/golang_toolkit/math/graph"
)

type Unit struct {
	eca.ECA
	Key    string                 `json:"key"`
	Pre    []*Unit                `json:"pre"`
	Post   []*Unit                `json:"post"`
	Params map[string]interface{} `json:"data"`
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
