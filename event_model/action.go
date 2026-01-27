package event_model

type Action interface {
	Execute() error
	GetAttributes() []string
}

type ActionBase struct {
	Attributes []string `json:"attributes"`
	Params     map[string]interface{}
}

func (c *ActionBase) GetAttributes() []string {
	return c.Attributes
}

func (c *ActionBase) Execute() error {
	return nil
}
