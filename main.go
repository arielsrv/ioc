package main

import (
	"context"
	"ioc/core"
	_ "ioc/docs"
	"log"

	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
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
// @BasePath /
func main() {
	app := fx.New(
		fx.Provide(core.NewHTTPClient),
		fx.Provide(core.NewPingService),
		fx.Provide(core.NewPingHandler),
		fx.Provide(core.NewApp),
		fx.Provide(logrus.New),
		fx.WithLogger(func(log *logrus.Logger) fxevent.Logger {
			return &fxevent.ConsoleLogger{
				W: log.Writer(),
			}
		}),
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
