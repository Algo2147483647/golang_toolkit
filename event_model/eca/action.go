package eca

import "context"

type Action interface {
	Execute(ctx context.Context) error
	GetAttributes(ctx context.Context) []string
}

type ActionBase struct {
	Attributes []string `json:"attributes"`
	Params     map[string]interface{}
}

func (c *ActionBase) Execute(ctx context.Context) error {
	return nil
}

func (c *ActionBase) GetAttributes(ctx context.Context) []string {
	return c.Attributes
}
