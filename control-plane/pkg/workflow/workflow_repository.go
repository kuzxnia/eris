package workflow

type WorkflowRepository interface {
	GetWorkflow(string) (*Workflow, error)
	CreateWorkflow(*Workflow) error
	UpdateWorkflow(*Workflow) error
}

type WorkflowConfigRepository interface {
	GetWorkflowConfig(string) (*WorkflowConfig, error)
	CreateWorkflowConfig(*WorkflowConfig) error
	UpdateWorkflowConfig(*WorkflowConfig) error
}
