package event_model

import "fmt"

type UnitGroup struct {
	Key    string
	Units  []*Unit                `json:"units"`
	Params map[string]interface{} `json:"params"`
}

type UnitGroupInstance struct {
	Key       string
	State     string
	Round     int64
	Units     []*Unit `json:"units"`
	UnitGroup *UnitGroup
	Params    map[string]interface{} `json:"params"`
}

const (
	UnitGroupStateNotStarted = "not_started"
	UnitGroupStateInProgress = "in_progress"
	UnitGroupStateCompleted  = "completed"
)

func (f *UnitGroupInstance) SetIdempotentKey(idempotentKey string) {
	f.Key = fmt.Sprintf("%s_%s_%d", idempotentKey, f.UnitGroup.Key, f.Round)
}
