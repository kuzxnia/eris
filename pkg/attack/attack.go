package attack

import "go.uber.org/fx"

type Attack interface {
	Perform()
	AttackType() string
}

func AsAttack(f any) fx.Option {
	return fx.Provide(
		fx.Annotate(
			f,
			fx.As(new(Attack)),
			fx.ResultTags(`group:"attacks"`),
		),
	)
}
