package event_model

type Policy struct {
	Key   string                 `json:"key"`
	Rules []*Rule                `json:"rules"`
	Data  map[string]interface{} `json:"data"`
}
