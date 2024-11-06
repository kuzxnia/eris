package cli

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"go.uber.org/fx"
)

func ProvideCli(commands []*cli.Command) *cli.App {
	app := &cli.App{
		Commands: commands,
	}
	return app
}

func RunCliApp(app *cli.App) {
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

var Module = fx.Options(
	fx.NopLogger, // disable fx logs in cli app

	AttackCliModule,
	fx.Provide(
		fx.Annotate(ProvideCli, fx.ParamTags(`group:"cli-command"`)),
	),
	fx.Invoke(RunCliApp),
)
