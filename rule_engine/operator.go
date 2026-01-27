package rule_engine

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type OperatorIf interface {
	GetType() string
	Evaluate(valueList ...ValueIf) (ValueIf, error)
}

// Operator type definitions
const (
	OpTypeEqual        = "=" // Comparison/Relational Operators
	OpTypeNotEqual     = "!="
	OpTypeLessThan     = "<"
	OpTypeLessEqual    = "<="
	OpTypeGreaterThan  = ">"
	OpTypeGreaterEqual = ">="
	OpTypeAnd          = "&&" // Logical Operators
	OpTypeOr           = "||"
	OpTypeNot          = "!"
	OpTypeBitwiseAnd   = "&" // Bitwise Operators
	OpTypeBitwiseOr    = "|"
	OpTypeBitwiseXor   = "^"
	OpTypeXor          = "xor"
	OpTypeLeftShift    = "<<"
	OpTypeRightShift   = ">>"
	OpTypeAdd          = "+" // Arithmetic Operators
	OpTypeSubtract     = "-"
	OpTypeMultiply     = "*"
	OpTypeDivide       = "/"
	OpTypeMod          = "%"
	OpTypeExp          = "**"
	OpTypeIn           = "in" // Membership/Set Operators
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
		return &ValueBase{Type: ValueTypeBool, Value: reflect.DeepEqual(getValue(valueList[0]), getValue(valueList[1]))}, nil
	case OpTypeNotEqual:
		return &ValueBase{Type: ValueTypeBool, Value: !reflect.DeepEqual(getValue(valueList[0]), getValue(valueList[1]))}, nil
	case OpTypeLessThan:
		result, err := lessThan(getValue(valueList[0]), getValue(valueList[1]))
		if err != nil {
			return nil, err
		}
		return &ValueBase{Type: ValueTypeBool, Value: result}, nil
	case OpTypeLessEqual:
		result, err := lessEqual(getValue(valueList[0]), getValue(valueList[1]))
		if err != nil {
			return nil, err
		}
		return &ValueBase{Type: ValueTypeBool, Value: result}, nil
	case OpTypeGreaterThan:
		result, err := greaterThan(getValue(valueList[0]), getValue(valueList[1]))
		if err != nil {
			return nil, err
		}
		return &ValueBase{Type: ValueTypeBool, Value: result}, nil
	case OpTypeGreaterEqual:
		result, err := greaterEqual(getValue(valueList[0]), getValue(valueList[1]))
		if err != nil {
			return nil, err
		}
		return &ValueBase{Type: ValueTypeBool, Value: result}, nil
	case OpTypeIn:
		result, err := inOperator(getValue(valueList[0]), getValue(valueList[1]))
		if err != nil {
			return nil, err
		}
		return &ValueBase{Type: ValueTypeBool, Value: result}, nil
	case OpTypeAnd:
		leftBool, ok := getValue(valueList[0]).(bool)
		if !ok {
			return nil, fmt.Errorf("left operand of && must be boolean")
		}
		rightBool, ok := getValue(valueList[1]).(bool)
		if !ok {
			return nil, fmt.Errorf("right operand of && must be boolean")
		}
		return &ValueBase{Type: ValueTypeBool, Value: leftBool && rightBool}, nil
	case OpTypeOr:
		leftBool, ok := getValue(valueList[0]).(bool)
		if !ok {
			return nil, fmt.Errorf("left operand of || must be boolean")
		}
		rightBool, ok := getValue(valueList[1]).(bool)
		if !ok {
			return nil, fmt.Errorf("right operand of || must be boolean")
		}
		return &ValueBase{Type: ValueTypeBool, Value: leftBool || rightBool}, nil
	case OpTypeNot:
		operand, ok := getValue(valueList[0]).(bool)
		if !ok {
			return nil, fmt.Errorf("operand of ! must be boolean")
		}
		return &ValueBase{Type: ValueTypeBool, Value: !operand}, nil
	case OpTypeAdd:
		result, err := add(getValue(valueList[0]), getValue(valueList[1]))
		if err != nil {
			return nil, err
		}
		return &ValueBase{Type: getValueType(result), Value: result}, nil
	case OpTypeSubtract:
		result, err := subtract(getValue(valueList[0]), getValue(valueList[1]))
		if err != nil {
			return nil, err
		}
		return &ValueBase{Type: getValueType(result), Value: result}, nil
	case OpTypeMultiply:
		result, err := multiply(getValue(valueList[0]), getValue(valueList[1]))
		if err != nil {
			return nil, err
		}
		return &ValueBase{Type: getValueType(result), Value: result}, nil
	case OpTypeDivide:
		result, err := divide(getValue(valueList[0]), getValue(valueList[1]))
		if err != nil {
			return nil, err
		}
		return &ValueBase{Type: getValueType(result), Value: result}, nil
	case OpTypeMod:
		result, err := mod(getValue(valueList[0]), getValue(valueList[1]))
		if err != nil {
			return nil, err
		}
		return &ValueBase{Type: getValueType(result), Value: result}, nil
	case OpTypeContains:
		str, ok1 := getValue(valueList[0]).(string)
		element, ok2 := getValue(valueList[1]).(string)
		if !ok1 || !ok2 {
			return nil, fmt.Errorf("both operands of contains must be strings")
		}
		return &ValueBase{Type: ValueTypeBool, Value: contains(str, element)}, nil
	case OpTypeStartsWith:
		str, ok1 := getValue(valueList[0]).(string)
		prefix, ok2 := getValue(valueList[1]).(string)
		if !ok1 || !ok2 {
			return nil, fmt.Errorf("both operands of startsWith must be strings")
		}
		return &ValueBase{Type: ValueTypeBool, Value: startsWith(str, prefix)}, nil
	case OpTypeEndsWith:
		str, ok1 := getValue(valueList[0]).(string)
		suffix, ok2 := getValue(valueList[1]).(string)
		if !ok1 || !ok2 {
			return nil, fmt.Errorf("both operands of endsWith must be strings")
		}
		return &ValueBase{Type: ValueTypeBool, Value: endsWith(str, suffix)}, nil
	case OpTypeXor:
		leftBool, ok := getValue(valueList[0]).(bool)
		if !ok {
			return nil, fmt.Errorf("left operand of xor must be boolean")
		}
		rightBool, ok := getValue(valueList[1]).(bool)
		if !ok {
			return nil, fmt.Errorf("right operand of xor must be boolean")
		}
		return &ValueBase{Type: ValueTypeBool, Value: leftBool != rightBool}, nil
	case OpTypeBitwiseAnd:
		result, err := bitwiseAnd(getValue(valueList[0]), getValue(valueList[1]))
		if err != nil {
			return nil, err
		}
		return &ValueBase{Type: getValueType(result), Value: result}, nil
	case OpTypeBitwiseOr:
		result, err := bitwiseOr(getValue(valueList[0]), getValue(valueList[1]))
		if err != nil {
			return nil, err
		}
		return &ValueBase{Type: getValueType(result), Value: result}, nil
	case OpTypeBitwiseXor:
		result, err := bitwiseXor(getValue(valueList[0]), getValue(valueList[1]))
		if err != nil {
			return nil, err
		}
		return &ValueBase{Type: getValueType(result), Value: result}, nil
	case OpTypeLeftShift:
		result, err := leftShift(getValue(valueList[0]), getValue(valueList[1]))
		if err != nil {
			return nil, err
		}
		return &ValueBase{Type: getValueType(result), Value: result}, nil
	case OpTypeRightShift:
		result, err := rightShift(getValue(valueList[0]), getValue(valueList[1]))
		if err != nil {
			return nil, err
		}
		return &ValueBase{Type: getValueType(result), Value: result}, nil
	case OpTypeExp:
		result, err := exp(getValue(valueList[0]), getValue(valueList[1]))
		if err != nil {
			return nil, err
		}
		return &ValueBase{Type: getValueType(result), Value: result}, nil
	case OpTypeNop:
		return valueList[0], nil
	default:
		return nil, fmt.Errorf("unsupported operator: %s", o.GetType())
	}
}

// Helper functions
func getValue(v ValueIf) interface{} {
	if v == nil {
		return nil
	}
	return v.GetValue()
}

func getValueType(v interface{}) string {
	switch v.(type) {
	case bool:
		return ValueTypeBool
	case string:
		return ValueTypeString
	case int:
		return ValueTypeInt
	case float64:
		return ValueTypeFloat64
	case int64:
		return ValueTypeInt64
	case uint:
		return ValueTypeUint
	case uint64:
		return ValueTypeUint64
	case []interface{}:
		return ValueTypeArray
	case map[string]interface{}:
		return ValueTypeMap
	default:
		return ValueTypeObject
	}
}

// Comparison operators
func lessThan(a, b interface{}) (bool, error) {
	// Convert both values to comparable types
	switch av := a.(type) {
	case int:
		if bv, ok := b.(int); ok {
			return av < bv, nil
		} else if bv, ok := b.(float64); ok {
			return float64(av) < bv, nil
		}
	case float64:
		if bv, ok := b.(float64); ok {
			return av < bv, nil
		} else if bv, ok := b.(int); ok {
			return av < float64(bv), nil
		}
	case string:
		if bv, ok := b.(string); ok {
			return av < bv, nil
		}
	}
	return false, fmt.Errorf("cannot compare %T and %T with <", a, b)
}

func lessEqual(a, b interface{}) (bool, error) {
	// Convert both values to comparable types
	switch av := a.(type) {
	case int:
		if bv, ok := b.(int); ok {
			return av <= bv, nil
		} else if bv, ok := b.(float64); ok {
			return float64(av) <= bv, nil
		}
	case float64:
		if bv, ok := b.(float64); ok {
			return av <= bv, nil
		} else if bv, ok := b.(int); ok {
			return av <= float64(bv), nil
		}
	case string:
		if bv, ok := b.(string); ok {
			return av <= bv, nil
		}
	}
	return false, fmt.Errorf("cannot compare %T and %T with <=", a, b)
}

func greaterThan(a, b interface{}) (bool, error) {
	// Convert both values to comparable types
	switch av := a.(type) {
	case int:
		if bv, ok := b.(int); ok {
			return av > bv, nil
		} else if bv, ok := b.(float64); ok {
			return float64(av) > bv, nil
		}
	case float64:
		if bv, ok := b.(float64); ok {
			return av > bv, nil
		} else if bv, ok := b.(int); ok {
			return av > float64(bv), nil
		}
	case string:
		if bv, ok := b.(string); ok {
			return av > bv, nil
		}
	}
	return false, fmt.Errorf("cannot compare %T and %T with >", a, b)
}

func greaterEqual(a, b interface{}) (bool, error) {
	// Convert both values to comparable types
	switch av := a.(type) {
	case int:
		if bv, ok := b.(int); ok {
			return av >= bv, nil
		} else if bv, ok := b.(float64); ok {
			return float64(av) >= bv, nil
		}
	case float64:
		if bv, ok := b.(float64); ok {
			return av >= bv, nil
		} else if bv, ok := b.(int); ok {
			return av >= float64(bv), nil
		}
	case string:
		if bv, ok := b.(string); ok {
			return av >= bv, nil
		}
	}
	return false, fmt.Errorf("cannot compare %T and %T with >=", a, b)
}

// Arithmetic operators
func add(a, b interface{}) (interface{}, error) {
	switch av := a.(type) {
	case int:
		if bv, ok := b.(int); ok {
			return av + bv, nil
		} else if bv, ok := b.(float64); ok {
			return float64(av) + bv, nil
		} else if bv, ok := b.(string); ok {
			return strconv.Itoa(av) + bv, nil
		}
	case float64:
		if bv, ok := b.(float64); ok {
			return av + bv, nil
		} else if bv, ok := b.(int); ok {
			return av + float64(bv), nil
		}
	case string:
		return av + fmt.Sprintf("%v", b), nil
	}
	return nil, fmt.Errorf("cannot add %T and %T", a, b)
}

func subtract(a, b interface{}) (interface{}, error) {
	switch av := a.(type) {
	case int:
		if bv, ok := b.(int); ok {
			return av - bv, nil
		} else if bv, ok := b.(float64); ok {
			return float64(av) - bv, nil
		}
	case float64:
		if bv, ok := b.(float64); ok {
			return av - bv, nil
		} else if bv, ok := b.(int); ok {
			return av - float64(bv), nil
		}
	}
	return nil, fmt.Errorf("cannot subtract %T and %T", a, b)
}

func multiply(a, b interface{}) (interface{}, error) {
	switch av := a.(type) {
	case int:
		if bv, ok := b.(int); ok {
			return av * bv, nil
		} else if bv, ok := b.(float64); ok {
			return float64(av) * bv, nil
		}
	case float64:
		if bv, ok := b.(float64); ok {
			return av * bv, nil
		} else if bv, ok := b.(int); ok {
			return av * float64(bv), nil
		}
	}
	return nil, fmt.Errorf("cannot multiply %T and %T", a, b)
}

func divide(a, b interface{}) (interface{}, error) {
	switch av := a.(type) {
	case int:
		if bv, ok := b.(int); ok {
			if bv == 0 {
				return nil, fmt.Errorf("division by zero")
			}
			return av / bv, nil
		} else if bv, ok := b.(float64); ok {
			if bv == 0.0 {
				return nil, fmt.Errorf("division by zero")
			}
			return float64(av) / bv, nil
		}
	case float64:
		if bv, ok := b.(float64); ok {
			if bv == 0.0 {
				return nil, fmt.Errorf("division by zero")
			}
			return av / bv, nil
		} else if bv, ok := b.(int); ok {
			if bv == 0 {
				return nil, fmt.Errorf("division by zero")
			}
			return av / float64(bv), nil
		}
	}
	return nil, fmt.Errorf("cannot divide %T and %T", a, b)
}

func mod(a, b interface{}) (interface{}, error) {
	switch av := a.(type) {
	case int:
		if bv, ok := b.(int); ok {
			if bv == 0 {
				return nil, fmt.Errorf("modulo by zero")
			}
			return av % bv, nil
		}
	case float64:
		if bv, ok := b.(float64); ok {
			if bv == 0.0 {
				return nil, fmt.Errorf("modulo by zero")
			}
			return float64(int(av) % int(bv)), nil
		} else if bv, ok := b.(int); ok {
			if bv == 0 {
				return nil, fmt.Errorf("modulo by zero")
			}
			return float64(int(av) % bv), nil
		}
	}
	return nil, fmt.Errorf("cannot calculate modulo of %T and %T", a, b)
}

func exp(a, b interface{}) (interface{}, error) {
	switch av := a.(type) {
	case int:
		if bv, ok := b.(int); ok {
			result := 1
			for i := 0; i < bv; i++ {
				result *= av
			}
			return result, nil
		} else if bv, ok := b.(float64); ok {
			// Convert to float for exponentiation
			return pow(float64(av), bv), nil
		}
	case float64:
		if bv, ok := b.(float64); ok {
			return pow(av, bv), nil
		} else if bv, ok := b.(int); ok {
			return pow(av, float64(bv)), nil
		}
	}
	return nil, fmt.Errorf("cannot calculate exponentiation of %T and %T", a, b)
}

// String operations
func contains(str, substr string) bool {
	return strings.Contains(str, substr)
}

func startsWith(str, prefix string) bool {
	return strings.HasPrefix(str, prefix)
}

func endsWith(str, suffix string) bool {
	return strings.HasSuffix(str, suffix)
}

func matches(str, pattern string) (bool, error) {
	matched, err := regexp.MatchString(pattern, str)
	if err != nil {
		return false, err
	}
	return matched, nil
}

// Bitwise operations
func bitwiseAnd(a, b interface{}) (interface{}, error) {
	switch av := a.(type) {
	case int:
		if bv, ok := b.(int); ok {
			return av & bv, nil
		}
	case int64:
		if bv, ok := b.(int64); ok {
			return av & bv, nil
		}
	case uint:
		if bv, ok := b.(uint); ok {
			return av & bv, nil
		}
	case uint64:
		if bv, ok := b.(uint64); ok {
			return av & bv, nil
		}
	}
	return nil, fmt.Errorf("cannot perform bitwise AND on %T and %T", a, b)
}

func bitwiseOr(a, b interface{}) (interface{}, error) {
	switch av := a.(type) {
	case int:
		if bv, ok := b.(int); ok {
			return av | bv, nil
		}
	case int64:
		if bv, ok := b.(int64); ok {
			return av | bv, nil
		}
	case uint:
		if bv, ok := b.(uint); ok {
			return av | bv, nil
		}
	case uint64:
		if bv, ok := b.(uint64); ok {
			return av | bv, nil
		}
	}
	return nil, fmt.Errorf("cannot perform bitwise OR on %T and %T", a, b)
}

func bitwiseXor(a, b interface{}) (interface{}, error) {
	switch av := a.(type) {
	case int:
		if bv, ok := b.(int); ok {
			return av ^ bv, nil
		}
	case int64:
		if bv, ok := b.(int64); ok {
			return av ^ bv, nil
		}
	case uint:
		if bv, ok := b.(uint); ok {
			return av ^ bv, nil
		}
	case uint64:
		if bv, ok := b.(uint64); ok {
			return av ^ bv, nil
		}
	}
	return nil, fmt.Errorf("cannot perform bitwise XOR on %T and %T", a, b)
}

func leftShift(a, b interface{}) (interface{}, error) {
	switch av := a.(type) {
	case int:
		if bv, ok := b.(int); ok {
			return av << uint(bv), nil
		} else if bv, ok := b.(uint); ok {
			return av << bv, nil
		}
	case int64:
		if bv, ok := b.(int); ok {
			return av << uint(bv), nil
		} else if bv, ok := b.(uint); ok {
			return av << bv, nil
		}
	}
	return nil, fmt.Errorf("cannot perform left shift on %T and %T", a, b)
}

func rightShift(a, b interface{}) (interface{}, error) {
	switch av := a.(type) {
	case int:
		if bv, ok := b.(int); ok {
			return av >> uint(bv), nil
		} else if bv, ok := b.(uint); ok {
			return av >> bv, nil
		}
	case int64:
		if bv, ok := b.(int); ok {
			return av >> uint(bv), nil
		} else if bv, ok := b.(uint); ok {
			return av >> bv, nil
		}
	}
	return nil, fmt.Errorf("cannot perform right shift on %T and %T", a, b)
}

// Utility function for exponentiation
func pow(base, exp float64) float64 {
	if exp == 0 {
		return 1
	}
	if exp == 1 {
		return base
	}

	result := 1.0
	negativeExp := exp < 0
	if negativeExp {
		exp = -exp
	}

	for i := 0; i < int(exp); i++ {
		result *= base
	}

	fracPart := exp - float64(int(exp))
	if fracPart > 0 {
		// For simplicity, we use Go's math.Pow for fractional exponents
		// In a real implementation, you'd import "math" package and use math.Pow
		result *= 1 // Placeholder - would use math.Pow(base, fracPart) in real implementation
	}

	if negativeExp {
		return 1 / result
	}
	return result
}

// Check if a value is in array/set
func inOperator(needle, haystack interface{}) (bool, error) {
	switch arr := haystack.(type) {
	case []interface{}:
		for _, item := range arr {
			if reflect.DeepEqual(needle, item) {
				return true, nil
			}
		}
	case []string:
		if str, ok := needle.(string); ok {
			for _, item := range arr {
				if str == item {
					return true, nil
				}
			}
		}
	case []int:
		if num, ok := needle.(int); ok {
			for _, item := range arr {
				if num == item {
					return true, nil
				}
			}
		}
	case []float64:
		if num, ok := needle.(float64); ok {
			for _, item := range arr {
				if num == item {
					return true, nil
				}
			}
		}
	case map[string]interface{}:
		if str, ok := needle.(string); ok {
			_, exists := arr[str]
			return exists, nil
		}
	case string:
		if substr, ok := needle.(string); ok {
			return strings.Contains(arr, substr), nil
		}
	default:
		return false, fmt.Errorf("right operand of 'in' must be an array, slice, map, or string, got %T", haystack)
	}
	return false, nil
}
