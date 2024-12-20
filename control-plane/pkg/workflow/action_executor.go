package workflow

import (
	"github.com/kuzxnia/eris/control-plane/pkg/agent"
	"github.com/kuzxnia/eris/control-plane/pkg/resource"
)

type ActionExecutor interface {
	Run(*ActionExecutionContext) (chan *ActionResult, error)
}

type AgentActionExecutor struct {
	resourceClient resource.ResourceClient
	repository     agent.AgentRepository
}

func ProvideAgentActionExecutor(resourceClient resource.ResourceClient, agentRepository agent.AgentRepository) *AgentActionExecutor {
	return &AgentActionExecutor{
		resourceClient: resourceClient,
		repository:     agentRepository,
	}
}

func (e *AgentActionExecutor) Run(action *Action, resultChannel chan *ActionResult) error {
	return nil
}
