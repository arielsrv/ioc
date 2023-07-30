package core

import "github.com/gofiber/fiber/v2"

type IPingHandler interface {
	Ping(ctx *fiber.Ctx) error
}

type PingHandler struct {
	pingService IPingService
}

func NewPingHandler(pingService IPingService) IPingHandler {
	return &PingHandler{
		pingService: pingService,
	}
}

func (p PingHandler) Ping(ctx *fiber.Ctx) error {
	message := p.pingService.Ping()
	return ctx.SendString(message)
}
