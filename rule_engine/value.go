package rule_engine

import (
	"reflect"
)

type ValueIf interface {
	GetType() string
	GetValue() interface{}
	SetValue(value interface{})
}

const (
	ValueTypeNull       = "null" // basic primitives
	ValueTypeBool       = "bool"
	ValueTypeInt8       = "int8" // Signed integers
	ValueTypeInt16      = "int16"
	ValueTypeInt32      = "int32"
	ValueTypeInt64      = "int64"
	ValueTypeUint8      = "uint8" // Unsigned integers
	ValueTypeUint16     = "uint16"
	ValueTypeUint32     = "uint32"
	ValueTypeUint64     = "uint64"
	ValueTypeFloat32    = "float32" // Floating point
	ValueTypeFloat64    = "float64"
	ValueTypeComplex64  = "complex64" // Complex numbers
	ValueTypeComplex128 = "complex128"
	ValueTypeByte       = "byte" // alias for uint8
	ValueTypeRune       = "rune" // alias for int32
	ValueTypeString     = "string"
	ValueTypeArray      = "array" // Composite ValueTypes
	ValueTypeSet        = "set"
	ValueTypeMap        = "map"
	ValueTypeStruct     = "struct"
	ValueTypeInterface  = "interface"
	ValueTypeFunc       = "func"
)

type ValueBase struct {
	Type  string
	ID    string
	Value interface{}
}

func (v *ValueBase) GetType() string {
	return v.Type
}

func (v *ValueBase) GetValue() interface{} {
	return v.Value
}

func (v *ValueBase) SetValue(value interface{}) {
	v.Value = value
}

// Helper function to determine the type of a value
func GetValueType(val interface{}) string {
	if val == nil {
		return ValueTypeNil
	}

	rv := reflect.ValueOf(val)
	switch rv.Kind() {
	case reflect.Bool:
		return ValueTypeBool
	case reflect.String:
		return ValueTypeString
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
		return ValueTypeInt
	case reflect.Int64:
		return ValueTypeInt64
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32:
		return ValueTypeUint
	case reflect.Uint64:
		return ValueTypeUint64
	case reflect.Float32, reflect.Float64:
		return ValueTypeFloat64
	case reflect.Array, reflect.Slice:
		return ValueTypeArray
	case reflect.Map:
		return ValueTypeMap
	default:
		return ValueTypeObject
	}
}

// Create a new ValueBase with automatic type detection
func NewValue(value interface{}) *ValueBase {
	return &ValueBase{
		Type:  GetValueType(value),
		Value: value,
	}
}
