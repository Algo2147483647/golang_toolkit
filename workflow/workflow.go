package workflow

import (
	"context"
)

type Workflow struct {
	Name      string
	Status    string
	WorkSteps []WorkStep
}

type WorkStepStatus string

const (
	WorkStepStatusPending     = "Pending"
	WorkStepStatusProgressing = "Progressing"
	WorkStepStatusCompleted   = "Completed"
	WorkStepStatusFailed      = "Failed"
)

func NewWorkflow(name string) *Workflow {
	return &Workflow{
		Name:   name,
		Status: WorkStepStatusPending,
	}
}

func (f *Workflow) Process(ctx context.Context, context interface{}) error {
	stepIndex := 0

	for stepIndex < len(f.WorkSteps) {

	}

	return nil
}
