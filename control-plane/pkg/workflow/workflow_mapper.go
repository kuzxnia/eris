package workflow

import (
	"github.com/kuzxnia/eris/control-plane/pkg/web/dto"
)

type WorkflowMapper struct{}

func ProvideWorkflowMapper() *WorkflowMapper {
	return &WorkflowMapper{}
}

func (m *WorkflowMapper) MapWorkflowConfig(request *dto.WorkflowConfigRequest) (*WorkflowConfig, error) {
	workflow := WorkflowConfig{
		Name:     request.Name,
		Contexts: Contexts(request.Contexts),
		Sources:  Sources(request.Sources),
		Actions:  make([]*Action, len(request.Actions)),
	}
	for i, action := range request.Actions {
		workflow.Actions[i] = &Action{
			Name:    action.Name,
			Type:    action.Type,
			Timeout: action.Timeout,
			Selector: &Selector{
				Type:                  action.Selector.Type,
				PodAffectedPercentage: action.Selector.PodAffectedPercentage,
				Label:                 action.Selector.Label,
				Namespaces:            action.Selector.Namespaces,
				ActionName:            action.Selector.ActionName,
				TargetPod:             action.Selector.TargetPod,
			},
			Probe: &Probe{
				Type:                   action.Probe.Type,
				Mode:                   action.Probe.Mode,
				Sources:                action.Probe.Sources,
				Check:                  action.Probe.Check,
				Frequency:              action.Probe.Frequency,
				Timeout:                action.Probe.Timeout,
				BlockAfterUntilSuccess: action.Probe.BlockAfterUntilSuccess,
				BlockTimeout:           action.Probe.BlockTimeout,
			},
			ProcessName:   action.ProcessName,
			ContainerName: action.ContainerName,
		}
	}

	return &workflow, nil
}
