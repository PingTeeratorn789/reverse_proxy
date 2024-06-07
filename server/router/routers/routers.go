package routers

import (
	"github.com/PingTeeratorn789/reverse_proxy/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRouters(server *fiber.App, handlers handlers.Handlers) {
	server.Get("/healthcheck", handlers.HealthHandler.HealthCheck)
}
