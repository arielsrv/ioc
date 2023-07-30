package core

import (
	"net/http"

	"github.com/gofiber/swagger"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type Container struct {
	fx.In

	PingHandler IPingHandler
}

func Routes(app *fiber.App, c Container) {
	app.Add(http.MethodGet, "/swagger/*", swagger.HandlerDefault)
	app.Add(http.MethodGet, "/ping", c.PingHandler.Ping)
}
