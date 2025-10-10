package rule_engine

import (
	"fmt"
	"reflect"
	"strings"
)

// Operator type definitions
const (
	OpEqual        = "="
	OpNotEqual     = "!="
	OpLessThan     = "<"
	OpLessEqual    = "<="
	OpGreaterThan  = ">"
	OpGreaterEqual = ">="
	OpAnd          = "&&"
	OpOr           = "||"
	OpNot          = "!"
	OpAdd          = "+"
	OpSubtract     = "-"
	OpMultiply     = "*"
	OpDivide       = "/"
	OpMod          = "%"
	OpIn           = "in"
)

// IsComparisonOperator checks if the operator is a comparison operator
func IsComparisonOperator(op string) bool {
	switch op {
	case OpEqual, OpNotEqual, OpLessThan, OpLessEqual, OpGreaterThan, OpGreaterEqual, OpIn:
		return true
	default:
		return false
	}
}

// IsLogicalOperator checks if the operator is a logical operator
func IsLogicalOperator(op string) bool {
	switch op {
	case OpAnd, OpOr, OpNot:
		return true
	default:
		return false
	}
}

// IsArithmeticOperator checks if the operator is an arithmetic operator
func IsArithmeticOperator(op string) bool {
	switch op {
	case OpAdd, OpSubtract, OpMultiply, OpDivide, OpMod:
		return true
	default:
		return false
	}
}

// EvaluateOperators evaluates two values with the given operator
func EvaluateOperators(left, right interface{}, operator string) (bool, error) {
	switch operator {
	case OpEqual:
		return reflect.DeepEqual(left, right), nil
	case OpNotEqual:
		return !reflect.DeepEqual(left, right), nil
	case OpLessThan:
	case OpLessEqual:
	case OpGreaterThan:
	case OpGreaterEqual:
	case OpIn:
		return inOperator(left, right)
	case OpAnd:
		leftBool, ok := left.(bool)
		if !ok {
			return false, fmt.Errorf("left operand of && must be boolean")
		}
		rightBool, ok := right.(bool)
		if !ok {
			return false, fmt.Errorf("right operand of && must be boolean")
		}
		return leftBool && rightBool, nil
	case OpOr:
		leftBool, ok := left.(bool)
		if !ok {
			return false, fmt.Errorf("left operand of || must be boolean")
		}
		rightBool, ok := right.(bool)
		if !ok {
			return false, fmt.Errorf("right operand of || must be boolean")
		}
		return leftBool || rightBool, nil
	default:
		return false, fmt.Errorf("unsupported operator: %s", operator)
	}
	return false, nil
}

// inOperator checks if left value is in right value (which should be a slice or array)
func inOperator(left, right interface{}) (bool, error) {
	rightValue := reflect.ValueOf(right)
	switch rightValue.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < rightValue.Len(); i++ {
			if reflect.DeepEqual(left, rightValue.Index(i).Interface()) {
				return true, nil
			}
		}
		return false, nil
	case reflect.String:
		leftStr, ok := left.(string)
		if !ok {
			return false, fmt.Errorf("'in' operator with string right operand requires left operand to be string")
		}
		rightStr := rightValue.String()
		return strings.Contains(rightStr, leftStr), nil
	default:
		return false, fmt.Errorf("'in' operator not supported for right operand type: %v", rightValue.Kind())
	}
}
