package event_model

type UnitGroup struct {
	Units []*Unit `json:"units"`
	Round int64
}
