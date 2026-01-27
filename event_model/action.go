package event_model

type Action interface {
	Execute() error
	GetAttributes() []string
}

type ActionBase struct {
	Attributes []string `json:"attributes"`
}

func (c *ActionBase) GetAttributes() []string {
	return c.Attributes
}
