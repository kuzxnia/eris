package infra

import (
	"github.com/kuzxnia/eris/control-plane/internal/infra/repository"
	"github.com/kuzxnia/eris/control-plane/pkg/agent"
	"github.com/kuzxnia/eris/control-plane/pkg/workflow"
	"go.uber.org/fx"
)

var Module = fx.Options(

	fx.Provide(
		fx.Annotate(repository.ProvideInMemoryAgentRepository, fx.As(new(agent.AgentRepository))),
	),
	fx.Provide(
		fx.Annotate(repository.ProvideInMemoryWorkflowRepository, fx.As(new(workflow.WorkflowRepository))),
	),
	fx.Provide(
		fx.Annotate(repository.ProvideInMemoryWorkflowConfigRepository, fx.As(new(workflow.WorkflowConfigRepository))),
	),
)
