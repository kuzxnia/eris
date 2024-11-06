package main

import (
	"context"
	"fmt"

	"github.com/kuzxnia/eris/eris-agent/internal/interfaces/cli"
	"github.com/kuzxnia/eris/eris-agent/pkg/attack"
	"go.uber.org/fx"
)

func main() {
	// setup configuration
	app := fx.New(
		attack.Module,
		cli.Module,
	)

	if err := app.Start(context.Background()); err != nil {
		fmt.Println("[Fx] START FAILED\t" + err.Error())
		return
	}
	if err := app.Stop(context.Background()); err != nil {
		fmt.Println("[Fx] STOP FAILED\t" + err.Error())
		return
	}
}
