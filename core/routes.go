package core

import (
	"github.com/gofiber/swagger"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type Container struct {
	fx.In

	MessageHandler PingHandler
}

func Routes(server *fiber.App, c Container) {
	server.Add(http.MethodGet, "/ping", c.MessageHandler.Ping)
	server.Get("/swagger/*", swagger.HandlerDefault)
}
