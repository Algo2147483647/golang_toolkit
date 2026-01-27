package rule_engine

type ValueIf interface {
	GetType() string
	GetValue() interface{}
	SetValue(value interface{})
}

const (
	ValueTypeString  = "string"
	ValueTypeFloat64 = "float64"
	ValueTypeBool    = "bool"
	ValueTypeArray   = "array"
	ValueTypeObject  = "object"
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
