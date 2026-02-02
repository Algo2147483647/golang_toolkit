package eca

import "github.com/Algo2147483647/golang_toolkit/rule_engine"

type Condition interface {
	IsPass() bool
	GetAttributes() []string
}

type ConditionBase struct {
	Attributes []string           `json:"attributes"`
	Node       rule_engine.NodeIf `json:"node"`
}

func (c *ConditionBase) IsPass() bool {
	result, err := c.Node.Evaluate()
	if err != nil {
		return false
	}
	return result.GetValue().(bool)
}

func (c *ConditionBase) GetAttributes() []string {
	return c.Attributes
}
