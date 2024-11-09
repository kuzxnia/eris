package agent

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(ProvideAgentMapper),
	fx.Provide(ProvideAgentService),
)
