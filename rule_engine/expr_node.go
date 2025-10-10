package rule_engine

import (
	"fmt"
)

// ExprNode is the base structure for expression nodes in the rule engine.
type ExprNode struct {
	Operator string      // The operator of the expression
	Value    interface{} // The value
	Left     *ExprNode   // Left operand
	Right    *ExprNode   // Right operand
}

// Evaluate evaluates the expression node with given context (variables)
func (n *ExprNode) Evaluate(context map[string]interface{}) (interface{}, error) {
	// If this is a leaf node (value node)
	if n.Operator == "" {
		// If the value is a variable (string starting with $)
		if variable, ok := n.Value.(string); ok && len(variable) > 0 && variable[0] == '$' {
			varName := variable[1:] // Remove the $
			if value, exists := context[varName]; exists {
				return value, nil
			}
			return nil, fmt.Errorf("variable not found in context: %s", varName)
		}
		// If it's a literal value
		return n.Value, nil
	}

	// Handle unary operators
	if n.Operator == OpNot {
		if n.Right == nil {
			return nil, fmt.Errorf("right operand required for operator: %s", n.Operator)
		}

		rightValue, err := n.Right.Evaluate(context)
		if err != nil {
			return nil, err
		}

		rightBool, ok := rightValue.(bool)
		if !ok {
			return nil, fmt.Errorf("operand of '!' must be boolean")
		}

		return !rightBool, nil
	}

	// Handle binary operators
	if n.Left == nil || n.Right == nil {
		return nil, fmt.Errorf("both operands required for operator: %s", n.Operator)
	}

	leftValue, err := n.Left.Evaluate(context)
	if err != nil {
		return nil, err
	}

	rightValue, err := n.Right.Evaluate(context)
	if err != nil {
		return nil, err
	}

	// Handle logical operators
	if IsLogicalOperator(n.Operator) {
		result, err := EvaluateOperators(leftValue, rightValue, n.Operator)
		if err != nil {
			return nil, err
		}
		return result, nil
	}

	// Handle comparison operators
	if IsComparisonOperator(n.Operator) {
		result, err := EvaluateOperators(leftValue, rightValue, n.Operator)
		if err != nil {
			return nil, err
		}
		return result, nil
	}

	// Handle arithmetic operators
	if IsArithmeticOperator(n.Operator) {
		return calculateArithmetic(leftValue, rightValue, n.Operator)
	}

	return nil, fmt.Errorf("unsupported operator: %s", n.Operator)
}

// calculateArithmetic performs arithmetic operations
func calculateArithmetic(left, right interface{}, operator string) (interface{}, error) {
	leftFloat, leftOK := toFloat64(left)
	rightFloat, rightOK := toFloat64(right)

	if !leftOK || !rightOK {
		return nil, fmt.Errorf("arithmetic operations require numeric operands")
	}

	switch operator {
	case OpAdd:
		return leftFloat + rightFloat, nil
	case OpSubtract:
		return leftFloat - rightFloat, nil
	case OpMultiply:
		return leftFloat * rightFloat, nil
	case OpDivide:
		if rightFloat == 0 {
			return nil, fmt.Errorf("division by zero")
		}
		return leftFloat / rightFloat, nil
	case OpMod:
		if rightFloat == 0 {
			return nil, fmt.Errorf("modulo by zero")
		}
		return float64(int64(leftFloat) % int64(rightFloat)), nil
	default:
		return nil, fmt.Errorf("unsupported arithmetic operator: %s", operator)
	}
}

// NewValueNode creates a new value node
func NewValueNode(value interface{}) *ExprNode {
	return &ExprNode{
		Value: value,
	}
}

// NewBinaryNode creates a new binary operation node
func NewBinaryNode(operator string, left, right *ExprNode) *ExprNode {
	return &ExprNode{
		Operator: operator,
		Left:     left,
		Right:    right,
	}
}

// NewUnaryNode creates a new unary operation node
func NewUnaryNode(operator string, right *ExprNode) *ExprNode {
	return &ExprNode{
		Operator: operator,
		Right:    right,
	}
}
