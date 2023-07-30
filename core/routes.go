package core

import (
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/gofiber/swagger"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type Container struct {
	fx.In

	Logger *logrus.Logger
}

type Handlers struct {
	fx.In

	PingHandler IPingHandler
}

func Routes(app *fiber.App, h Handlers) {
	app.Add(http.MethodGet, "/swagger/*", swagger.HandlerDefault)
	app.Add(http.MethodGet, "/ping", h.PingHandler.Ping)
}
