package handlers

import (
	"github.com/PingTeeratorn789/reverse_proxy/configs"
	"github.com/PingTeeratorn789/reverse_proxy/internal/handlers/health"

	"github.com/gofiber/fiber/v2"
)

type HealthHandlers interface {
	HealthCheck(ctx *fiber.Ctx) error
}

type Dependencies struct {
	Config configs.Config
}

type Handlers struct {
	HealthHandler HealthHandlers
}

func NewHandlers(d Dependencies) *Handlers {
	return &Handlers{
		HealthHandler: health.NewHandlers(health.Dependencies{
			Prefix:      d.Config.App.Prefix,
			ServiceName: d.Config.App.Name,
		}),
	}
}
