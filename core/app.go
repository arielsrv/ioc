package core

import (
	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type App struct {
	*fiber.App
}

func NewApp(container Container, handlers Handlers) *App {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		EnablePrintRoutes:     false,
	})

	prometheus := fiberprometheus.New("ioc")
	prometheus.RegisterAt(app, "/metrics")
	app.Use(prometheus.Middleware)

	app.Use(logger.New(logger.Config{
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
		Output: container.Logger.Writer(),
	}))

	Routes(app, handlers)

	return &App{
		App: app,
	}
}
