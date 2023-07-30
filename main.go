package main

import (
	"context"
	"ioc/core"
	"log"

	_ "ioc/docs"

	"go.uber.org/fx"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /.
func main() {
	app := fx.New(
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
