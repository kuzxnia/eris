package main

import (
	"fmt"

	"github.com/kuzxnia/eris/pkg/attack"
	"go.uber.org/fx"
)

func main() {
	// setup configuration
	fx.New(
		attack.Module,
		fx.Invoke(func(registry *attack.AttackRegistry) {
			fmt.Println("hello")
		}),
	).Run()
}
