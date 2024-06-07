package health

import (
	"os"

	"github.com/PingTeeratorn789/reverse_proxy/internal/handlers/common"
	"github.com/gofiber/fiber/v2"
)

type presenter struct {
	prefix      string
	serviceName string
}

func (p *presenter) toHealthResponse(ctx *fiber.Ctx) error {
	var (
		hostname, _ = os.Hostname()
		data        = fiber.Map{
			"message":  "alive",
			"hostname": hostname,
		}
		response = common.NewSuccessResponse(data, p.prefix, fiber.StatusOK)
	)
	return ctx.Status(fiber.StatusOK).JSON(response)
}
