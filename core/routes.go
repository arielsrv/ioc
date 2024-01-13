package core

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/sirupsen/logrus"
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
