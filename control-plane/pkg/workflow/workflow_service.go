package workflow

import (
	"fmt"

	"github.com/kuzxnia/eris/control-plane/pkg/web/dto"
)

type WorkflowService struct {
	mapper     *WorkflowMapper
	repository WorkflowRepository
	// agentService *agent.AgentService
}

func ProvideWorkflowService(
	// agentService *agent.AgentService,
	mapper *WorkflowMapper,
	repository WorkflowRepository,
) *WorkflowService {
	return &WorkflowService{
		mapper:     mapper,
		repository: repository,
	}
}

func (s *WorkflowService) CreateWorkflow(workflowRequest dto.WorkflowRequest) error {
	workflow, _ := s.mapper.Map(&workflowRequest)
	return s.repository.SaveWorkflow(workflow)
}

func (s *WorkflowService) RunWorkflow(workflow_name string) error {
	workflow, err := s.repository.GetWorkflow(workflow_name)
	if err != nil {
		return err
	}
	//
	fmt.Println(workflow)
	return nil
}
