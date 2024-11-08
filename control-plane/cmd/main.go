package main

import (
	app "github.com/kuzxnia/eris/control-plane/pkg"
	"go.uber.org/fx"
)

func main() {
	// setup configuration
	fx.New(
		app.Modules,
	).Run()
}
