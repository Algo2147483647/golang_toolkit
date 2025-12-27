package event_model

type Policy struct {
	Rules []*Rule `json:"rules"`
}
