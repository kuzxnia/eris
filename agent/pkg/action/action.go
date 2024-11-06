package action

import (
	"github.com/kuzxnia/eris/agent/pkg/action/k8s/pod"
	"go.uber.org/fx"
)

type Action interface {
	Perform() error
	ActionType() string
}

func AsAction(f any) fx.Option {
	// todo: double registration problem, blocked until https://github.com/uber-go/fx/issues/1214
	return fx.Provide(
		f,
		fx.Annotate(
			f, fx.As(new(Action)), fx.ResultTags(`group:"actions"`),
		),
	)
}

var Module = fx.Options(
	// k8s pod
	AsAction(pod.ProvideContainerKillAction),
	// AsAction(pod.ProvideCpuHogAction),
	// AsAction(pod.ProvideDiskFillAction),
	// AsAction(pod.ProvideIpSwapAction),
	// AsAction(pod.ProvideProcessFloodAction),
	// k8s node
	// AsAction(node.ProvideNodeDrainAction),
	// AsAction(node.ProvideNodeIpOverflowAction),
	// AsAction(node.ProvideNodeRestartAction),
	// AsAction(node.ProvideNodeSchedulingDisabledAction),

	fx.Provide(
		fx.Annotate(ProvideActionRegistry, fx.ParamTags(`group:"actions"`)),
	),
)
