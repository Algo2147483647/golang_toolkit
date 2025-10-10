package workflow

import "context"

// WorkStep represents a single step or stage in a workflow pipeline.
type WorkStep interface {
	GetName() string
	Process(ctx context.Context, context interface{}) (string, error)
}
