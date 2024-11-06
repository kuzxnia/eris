package attack

import (
	"github.com/kuzxnia/eris/eris-agent/pkg/attack/k8s/node"
	"github.com/kuzxnia/eris/eris-agent/pkg/attack/k8s/pod"
	"go.uber.org/fx"
)

type Attack interface {
	Perform()
	AttackType() string
}

func AsAttack(f any) fx.Option {
	// todo: double registration problem, blocked until https://github.com/uber-go/fx/issues/1214
	return fx.Provide(
		f,
		fx.Annotate(
			f, fx.As(new(Attack)), fx.ResultTags(`group:"attackCommands"`),
		),
	)
}

var Module = fx.Options(
	// k8s pod
	fx.Provide(pod.ProvideContainerKillAttack),
	fx.Provide(pod.ProvideCpuHogAttack),
	fx.Provide(pod.ProvideDiskFillAttack),
	fx.Provide(pod.ProvideIpSwapAttack),
	fx.Provide(pod.ProvideProcessFloodAttack),
	// k8s node
	fx.Provide(node.ProvideNodeDrainAttack),
	fx.Provide(node.ProvideNodeIpOverflowAttack),
	fx.Provide(node.ProvideNodeRestartAttack),
	fx.Provide(node.ProvideNodeSchedulingDisabledAttack),

	fx.Provide(
		fx.Annotate(ProvideAttackRegistry, fx.ParamTags(`group:"attackCommands"`)),
	),
)
