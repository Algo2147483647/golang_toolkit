package eca

import (
	"context"
	"github.com/Algo2147483647/golang_toolkit/rule_engine"
)

type Condition interface {
	IsPass(ctx context.Context) bool
	GetAttributes(ctx context.Context) []string
}

type ConditionBase struct {
	Attributes []string           `json:"attributes"`
	Node       rule_engine.NodeIf `json:"node"`
}

func (c *ConditionBase) IsPass(ctx context.Context) bool {
	result, err := c.Node.Evaluate()
	if err != nil {
		return false
	}
	return result.GetValue().(bool)
}

func (c *ConditionBase) GetAttributes(ctx context.Context) []string {
	return c.Attributes
}
