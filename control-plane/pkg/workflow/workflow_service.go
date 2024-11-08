package workflow

import "github.com/kuzxnia/eris/control-plane/pkg/interfaces/web/dto"

type WorkflowService struct {
	mapper *WorkflowMapper
}

func ProvideWorkflowService(mapper *WorkflowMapper) *WorkflowService {
	return &WorkflowService{
		mapper: mapper,
	}
}

func (s *WorkflowService) CreateWorkflow(workflowRequest dto.WorkflowRequest) error {
	// todo: check if such chat exists
	return nil
}
