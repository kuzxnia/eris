package attack

import (
	"github.com/kuzxnia/eris/pkg/attack/k8s/node"
	"github.com/kuzxnia/eris/pkg/attack/k8s/pod"
	"go.uber.org/fx"
)

var Module = fx.Options(
	// k8s pod
	AsAttack(pod.ProvideContainerKillAttack),
	AsAttack(pod.ProvideCpuHogAttack),
	AsAttack(pod.ProvideDiskFillAttack),
	AsAttack(pod.ProvideIpSwapAttack),
	AsAttack(pod.ProvideProcessFloodAttack),
	// k8s node
	AsAttack(node.ProvideNodeDrainAttack),
	AsAttack(node.ProvideNodeIpOverflowAttack),
	AsAttack(node.ProvideNodeRestartAttack),
	AsAttack(node.ProvideNodeSchedulingDisabledAttack),

	fx.Provide(
		fx.Annotate(
			ProvideAttackRegistry,
			fx.ParamTags(`group:"attacks"`),
		),
	),
)
