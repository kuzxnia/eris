package repository

import (
	"github.com/google/uuid"
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

func (r *InMemoryWorkflowRepository) CreateWorkflow(workflow *workflow.Workflow) error {
	workflow.Id = uuid.NewString()
	_, isPresent := r.workflows[workflow.Id]
	if isPresent {
		return exception.ResourceAlreadyExistsError{}
	}

	r.workflows[workflow.Id] = workflow
	return nil
}

func (r *InMemoryWorkflowRepository) UpdateWorkflow(workflow *workflow.Workflow) error {
	r.workflows[workflow.Id] = workflow
	return nil
}

func (r *InMemoryWorkflowRepository) GetWorkflow(workflowName string) (*workflow.Workflow, error) {
	workflow, isPresent := r.workflows[workflowName]
	if !isPresent {
		return nil, exception.NotFoundError{}
	}

	return workflow, nil
}
