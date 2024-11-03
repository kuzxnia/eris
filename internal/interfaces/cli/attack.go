package cli

import (
	"github.com/kuzxnia/eris/internal/infrastructure/k8s"
	"github.com/kuzxnia/eris/pkg/attack/k8s/pod"
	"github.com/urfave/cli/v2"
	"go.uber.org/fx"
)

var AttackCliModule = fx.Options(
	k8s.Module,
	// pod
	AsCliAttackCommand(ProvideContainerKillCommand),
	AsCliAttackCommand(ProvideCpuHogCommand),
	// node

	fx.Provide(
		fx.Annotate(
			ProvideAttackCli,
			fx.ParamTags(`group:"cli-attack-command"`),
			fx.ResultTags(`group:"cli-command"`),
		),
	),
)

func AsCliAttackCommand(f any) fx.Option {
	return fx.Provide(
		fx.Annotate(f, fx.ResultTags(`group:"cli-attack-command"`)),
	)
}

func ProvideAttackCli(commands []*cli.Command) *cli.Command {
	return &cli.Command{
		Name:        "attack",
		Aliases:     []string{"a"},
		Usage:       "List of chaos attack actions.",
		Subcommands: commands,
	}
}

func ProvideContainerKillCommand(attack *pod.ContainerKillAttack) *cli.Command {
	return &cli.Command{
		Name:  "kill-container",
		Usage: "kill-container <regex>",
		Action: func(cCtx *cli.Context) error {
			attack.Perform(nil)
			return nil
		},
	}
}

func ProvideCpuHogCommand(attack *pod.CpuHogAttack) *cli.Command {
	return &cli.Command{
		Name:  "cpu-hog",
		Usage: "cpu-hog <regex>",
		Action: func(cCtx *cli.Context) error {
			attack.Perform()
			return nil
		},
	}
}
