package event_model

type Rule struct {
	Events     []Event     `json:"event"`
	Conditions []Condition `json:"condition"`
	Actions    []Action    `json:"action"`
	Pre        []*Rule     `json:"pre"`
	Post       []*Rule     `json:"post"`
}
