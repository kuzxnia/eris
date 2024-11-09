package workflow

type WorkflowRepository interface {
	GetWorkflow(string) (*Workflow, error)
	SaveWorkflow(*Workflow) error
}
