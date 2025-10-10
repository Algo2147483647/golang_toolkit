package event_model

type PolicyGroup struct {
	Policies []*Policy `json:"policies"`
}
