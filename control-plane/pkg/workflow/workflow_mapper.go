package workflow

type WorkflowMapper struct{}

func ProvideWorkflowMapper() *WorkflowMapper {
	return &WorkflowMapper{}
}
