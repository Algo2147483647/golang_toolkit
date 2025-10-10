package workflow

type Management struct {
	workflows map[string]*Workflow
}

func NewManagement() *Management {
	return &Management{
		workflows: make(map[string]*Workflow),
	}
}

func (m *Management) RegisterWorkflow(key string, workflow *Workflow) *Management {
	m.workflows[key] = workflow
	return m
}

func (m *Management) GetWorkflow(name string) *Workflow {
	workflow, exists := m.workflows[name]
	if !exists {
		return nil
	}

	return workflow
}

func (m *Management) DeleteWorkflow(key string) *Management {
	if _, exists := m.workflows[key]; !exists {
		return m
	}
	delete(m.workflows, key)
	return m
}
