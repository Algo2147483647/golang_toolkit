package rule_engine

import (
	"fmt"
	"reflect"
)

type OperatorIf interface {
	GetType() string
	Evaluate(valueList ...ValueIf) (ValueIf, error)
}

// Operator type definitions
const (
	OpTypeEqual        = "="
	OpTypeNotEqual     = "!="
	OpTypeLessThan     = "<"
	OpTypeLessEqual    = "<="
	OpTypeGreaterThan  = ">"
	OpTypeGreaterEqual = ">="
	OpTypeAnd          = "&&"
	OpTypeOr           = "||"
	OpTypeNot          = "!"
	OpTypeAdd          = "+"
	OpTypeSubtract     = "-"
	OpTypeMultiply     = "*"
	OpTypeDivide       = "/"
	OpTypeMod          = "%"
	OpTypeIn           = "in"
)

type OperatorBase struct {
	Type string
}

func NewOperatorBase(operator string) *OperatorBase {
	return &OperatorBase{
		Type: operator,
	}
}

func (o *OperatorBase) GetType() string {
	return o.Type
}
func (o *OperatorBase) Evaluate(valueList ...ValueIf) (ValueIf, error) {
	switch o.GetType() {
	case OpTypeEqual:
		return reflect.DeepEqual(valueList[0], valueList[1]), nil
	case OpTypeNotEqual:
		return !reflect.DeepEqual(valueList[0], valueList[1]), nil
	case OpTypeLessThan:
	case OpTypeLessEqual:
	case OpTypeGreaterThan:
	case OpTypeGreaterEqual:
	case OpTypeIn:
		return inOperator(valueList[0], valueList[1])
	case OpTypeAnd:
		leftBool, ok := valueList[0].(bool)
		if !ok {
			return false, fmt.Errorf("left operand of && must be boolean")
		}
		rightBool, ok := valueList[1].(bool)
		if !ok {
			return false, fmt.Errorf("right operand of && must be boolean")
		}
		return leftBool && rightBool, nil
	case OpTypeOr:
		leftBool, ok := valueList[0].(bool)
		if !ok {
			return false, fmt.Errorf("left operand of || must be boolean")
		}
		rightBool, ok := valueList[1].(bool)
		if !ok {
			return false, fmt.Errorf("right operand of || must be boolean")
		}
		return leftBool || rightBool, nil
	default:
		return false, fmt.Errorf("unsupported operator: %s", operator)
	}
	return false, nil
}
