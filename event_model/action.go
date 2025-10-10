package event_model

type Action interface {
	Execute() error
}
