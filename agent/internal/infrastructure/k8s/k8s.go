package k8s

import (
	"github.com/kuzxnia/eris/agent/pkg/action/k8s"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		fx.Annotate(
			ProvideCompositeResourceProvider,
			fx.As(new(k8s.ResourceProvider)),
		),
	),
)
