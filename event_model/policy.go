package event_model

type Policy struct {
	PolicyBase
	Pre  *Policy `json:"pre"`
	Post *Policy `json:"post"`
}

type PolicyBase struct {
	Events     []Event     `json:"event"`
	Conditions []Condition `json:"condition"`
	Actions    []Action    `json:"action"`
}
