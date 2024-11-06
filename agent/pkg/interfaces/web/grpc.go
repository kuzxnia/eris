package web

import (
	"context"
	"fmt"
	"net"

	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Controller interface {
	Register(grpc.ServiceRegistrar) error
}

func NewGRPCServer(controllers []Controller, lc fx.Lifecycle) *grpc.Server {
	server := grpc.NewServer()

	for _, controller := range controllers {
		err := controller.Register(server)
		if err != nil {
			panic(err)
		}
	}

	reflection.Register(server)

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			// todo: provide custom port
			address := "0.0.0.0:" + "1234"
			// log.Info("Started lbot-agent on " + address)
			tcpListener, err := net.Listen("tcp", address)
			if err != nil {
				// log.Fatal("listen error:", err)
				return err
			}
			fmt.Println("before")

			go server.Serve(tcpListener)
			// log.Fatalf("failed to serve: %s", err)
			return nil
		},
		OnStop: func(context.Context) error {
			fmt.Println("server stop")
			server.GracefulStop()
			return nil
		},
	})

	return server
}

var Module = fx.Options(
	AsController(ProvideWorkflowController),

	fx.Provide(
		fx.Annotate(
			NewGRPCServer,
			fx.ParamTags(`group:"controllers"`),
		),
	),
	fx.Invoke(func(*grpc.Server) {}),
)

func AsController(f any) fx.Option {
	return fx.Provide(
		fx.Annotate(f, fx.As(new(Controller)), fx.ResultTags(`group:"controllers"`)),
	)
}
