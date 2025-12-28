package event_model

type Condition interface {
	Check() bool
}
