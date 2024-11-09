package web

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"github.com/kuzxnia/eris/control-plane/pkg/config"
	"github.com/kuzxnia/eris/control-plane/pkg/exception"
	httpException "github.com/kuzxnia/eris/control-plane/pkg/interfaces/web/exception"
	v1 "github.com/kuzxnia/eris/control-plane/pkg/interfaces/web/v1"
	"go.uber.org/fx"
)

func ProvideHttpServer(controllers []Controller, config *config.Config, lc fx.Lifecycle) *fiber.App {
	cfg := fiber.Config{
		ErrorHandler: httpException.ErrorHandler,
	}

	app := fiber.New(cfg)
	app.Use(
		recover.New(recover.Config{EnableStackTrace: true}),
	)

	app.Use(cors.New())

	api := app.Group("/api")
	v1 := api.Group("/v1")

	for _, controller := range controllers {
		controller.RegisterRoutes(v1)
	}

	api.Get("/swagger/*", swagger.HandlerDefault)

	lc.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				go func() {
					err := app.Listen(fmt.Sprintf("0.0.0.0:%d", config.Server.Port))
					exception.PanicLogging(err)
				}()

				return nil
			},
			OnStop: func(context.Context) error {
				return nil
			},
		},
	)

	return app
}

var Module = fx.Options(
	AsController(v1.ProvideHealthCheckController),

	AsController(v1.ProvideWorkflowController),

	fx.Provide(
		fx.Annotate(
			ProvideHttpServer,
			fx.ParamTags(`group:"controllers"`),
		),
	),
	fx.Invoke(func(*fiber.App) {}),
)

func AsController(f any) fx.Option {
	return fx.Provide(
		fx.Annotate(f, fx.As(new(Controller)), fx.ResultTags(`group:"controllers"`)),
	)
}
