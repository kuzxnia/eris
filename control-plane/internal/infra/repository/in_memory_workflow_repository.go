package repository

import (
	"github.com/kuzxnia/eris/control-plane/pkg/exception"
	"github.com/kuzxnia/eris/control-plane/pkg/workflow"
)

type InMemoryWorkflowRepository struct {
	workflows map[string]*workflow.Workflow
}

func ProvideInMemoryWorkflowRepository() *InMemoryWorkflowRepository {
	return &InMemoryWorkflowRepository{
		workflows: map[string]*workflow.Workflow{},
	}
}

func (r *InMemoryWorkflowRepository) SaveWorkflow(workflow *workflow.Workflow) error {
	_, isPresent := r.workflows[workflow.Name]
	if isPresent {
		return exception.ResourceAlreadyExistsError{}
	}

	r.workflows[workflow.Name] = workflow
	return nil
}

func (r *InMemoryWorkflowRepository) GetWorkflow(workflowName string) (*workflow.Workflow, error) {
	workflow, isPresent := r.workflows[workflowName]
	if !isPresent {
		return nil, exception.NotFoundError{}
	}

	return workflow, nil
}
