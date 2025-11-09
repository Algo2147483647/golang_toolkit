package rule_engine

type Value interface {
}

const (
	ValueTypeString = iota
	ValueTypeInt
	ValueTypeFloat
	ValueTypeBool
	ValueTypeArray
	ValueTypeObject
)

type ValueBase struct {
	Type  int
	Value interface{}
}

func NewValueBase(value interface{}) *ValueBase {
	return &ValueBase{}
}
