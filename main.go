package main

import (
	"context"
	"go.uber.org/fx"
	"ioc/core"
	"log"
)

func main() {
	app := fx.New(
		// fx.NopLogger,
		fx.Provide(core.NewHTTPClient),
		fx.Provide(core.NewPingService),
		fx.Provide(core.NewPingHandler),
		fx.Provide(core.NewApp),
		fx.Invoke(func(app *core.App) {
			err := app.Listen(":8080")
			if err != nil {
				log.Fatal(err)
			}
		}),
	)

	err := app.Start(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
