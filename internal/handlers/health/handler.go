package health

import (
	"github.com/gofiber/fiber/v2"
)

type Dependencies struct {
	Prefix      string
	ServiceName string
}
type handlers struct {
	presenter *presenter
}

func NewHandlers(d Dependencies) *handlers {
	return &handlers{
		presenter: &presenter{
			prefix:      d.Prefix,
			serviceName: d.ServiceName,
		},
	}
}

func (h *handlers) HealthCheck(ctx *fiber.Ctx) error {
	return h.presenter.toHealthResponse(ctx)
}
