package event_model

import "github.com/Algo2147483647/golang_toolkit/rule_engine"

type Condition interface {
	IsPass() bool
}

type ConditionBase struct {
	Node rule_engine.NodeIf `json:"node"`
}

func (c *ConditionBase) IsPass() bool {
	result, err := c.Node.Evaluate()
	if err != nil {
		return false
	}
	return result.GetValue().(bool)
}
