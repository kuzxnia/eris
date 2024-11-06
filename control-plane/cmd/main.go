package main 

import (
	"context"
	"fmt"

	"go.uber.org/fx"
)

func main() {
	// setup configuration
	app := fx.New()

	if err := app.Start(context.Background()); err != nil {
		fmt.Println("[Fx] START FAILED\t" + err.Error())
		return
	}
	if err := app.Stop(context.Background()); err != nil {
		fmt.Println("[Fx] STOP FAILED\t" + err.Error())
		return
	}
}
