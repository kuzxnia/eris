package main

import (
	"github.com/kuzxnia/eris/agent/pkg/action"
	"github.com/kuzxnia/eris/agent/pkg/interfaces/web"
	"go.uber.org/fx"
)

func main() {
	// setup configuration
	fx.New(
		action.Module,
		web.Module,
	).Run()
}
