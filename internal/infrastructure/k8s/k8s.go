package k8s

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		fx.Annotate(
			ProvideCompositeResourceProvider,
			fx.As(new(ResourceProvider)),
		),
	),
)
