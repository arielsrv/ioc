package core

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"net/http"
)

type Container struct {
	fx.In

	MessageHandler IPingHandler
}

func Routes(server *fiber.App, c Container) {
	server.Add(http.MethodGet, "/ping", c.MessageHandler.Ping)
}
