package event_model

type Rule struct {
	ECA
	Key    string                 `json:"key"`
	Pre    []*Rule                `json:"pre"`
	Post   []*Rule                `json:"post"`
	Policy *Policy                `json:"policy"`
	Data   map[string]interface{} `json:"data"`
}
