package workflow

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/kuzxnia/eris/control-plane/pkg/resource"
	"github.com/kuzxnia/eris/control-plane/pkg/web/dto"
	"github.com/samber/lo"
)

type WorkflowService struct {
	mapper                   *WorkflowMapper
	workflowRepository       WorkflowRepository
	workflowConfigRepository WorkflowConfigRepository
	selectorRegistry         *resource.ResourceSelectorRegistry
	actionExecutor           ActionExecutor
}

func ProvideWorkflowService(
	mapper *WorkflowMapper,
	workflowRepository WorkflowRepository,
	workflowConfigRepository WorkflowConfigRepository,
	selectorRegistry *resource.ResourceSelectorRegistry,
	actionExecutor ActionExecutor,
) *WorkflowService {
	return &WorkflowService{
		mapper:                   mapper,
		workflowRepository:       workflowRepository,
		workflowConfigRepository: workflowConfigRepository,
		selectorRegistry:         selectorRegistry,
		actionExecutor:           actionExecutor,
	}
}

// todo: this could be seperate config service
func (s *WorkflowService) ConfigureWorkflow(workflowConfigRequest dto.WorkflowConfigRequest) error {
	cfg, _ := s.mapper.MapWorkflowConfig(&workflowConfigRequest)

	return s.workflowConfigRepository.CreateWorkflowConfig(cfg)
}

func (s *WorkflowService) ReConfigureWorkflow(workflowConfigRequest dto.WorkflowConfigRequest) error {
	oldCfg, err := s.workflowConfigRepository.GetWorkflowConfig(workflowConfigRequest.Name)
	if err != nil {
		return err
	}

	cfg, _ := s.mapper.MapWorkflowConfig(&workflowConfigRequest)
	cfg.Version = oldCfg.Version + 1

	return s.workflowConfigRepository.CreateWorkflowConfig(cfg)
}

func (s *WorkflowService) CreateWorkflowFromConfig(workflowName string) (*Workflow, error) {
	workflowConfig, err := s.workflowConfigRepository.GetWorkflowConfig(workflowName)
	if err != nil {
		return nil, err
	}

	workflow := Workflow{
		Id:               uuid.NewString(),
		Status:           "active", // todo: fix, add statuse
		StartTs:          time.Now().UTC(),
		LastTransitionTs: time.Now().UTC(),
		Config:           workflowConfig,
	}

	err = s.workflowRepository.CreateWorkflow(&workflow)
	if err != nil {
		return nil, err
	}
	return &workflow, nil
}

// todo: could be seperate handler
func (s *WorkflowService) StartWorkflow(workflowName string) error {
	workflow, err := s.CreateWorkflowFromConfig(workflowName)
	if err != nil {
		return err
	}

	for _, action := range workflow.Config.Actions {
		resources, err := s.selectorRegistry.ResolveResourceSelector((*resource.Selector)(action.Selector))
		if err != nil {
			return err
		}
		action := ActionExecutionContext{
			Action:           action,
			Status:           "active",
			StartTs:          time.Now().UTC(),
			LastTransitionTs: time.Now().UTC(),
			Resources:        lo.Map(resources, func(item *resource.Resource, index int) *Resource { return &Resource{item.Name} }),
		}
		workflow.Actions = append(workflow.Actions, &action)

		if err = s.workflowRepository.UpdateWorkflow(workflow); err != nil {
			return err
		}

		resultChannel, err := s.actionExecutor.Run(&action)
		if err != nil {
			return err
		}

		// todo: add non blocking channel read
		for {
			result, ok := <-resultChannel

			if !ok {
				break
			}
			if result.Err != nil {
				// todo: handle compensation
				return result.Err
			}
			// todo: handle messages, append to workflow
			fmt.Println(result.Message)
		}
		close(resultChannel)
		// when to go to next action ??
		// when result chann will be closed
	}
	return nil
}
