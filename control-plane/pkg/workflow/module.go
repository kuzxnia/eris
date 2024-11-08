package workflow

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(ProvideWorkflowMapper),
	fx.Provide(ProvideWorkflowService),
)
