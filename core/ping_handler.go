package core

import (
	"github.com/gofiber/fiber/v2"
)

type IPingHandler interface {
	Ping(ctx *fiber.Ctx) error
}

type PingHandler struct {
	PingService IPingService
}

func NewPingHandler(pingService IPingService) IPingHandler {
	return &PingHandler{
		PingService: pingService,
	}
}

// Ping godoc
// @Summary	Check if the instance is online
// @Description	Ping
// @Tags	Check
// @Success	200
// @Produce	plain
// @Success	200	{string}	string	"pong"
// @Router		/ping [get].
func (p PingHandler) Ping(ctx *fiber.Ctx) error {
	message := p.PingService.Ping()
	return ctx.SendString(message)
}
