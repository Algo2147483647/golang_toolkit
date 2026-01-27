package rule_engine

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
