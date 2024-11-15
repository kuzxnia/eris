package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kuzxnia/eris/control-plane/pkg/web/dto"
)

func ProvideHealthCheckController() *HealthCheckController {
	return &HealthCheckController{}
}

type HealthCheckController struct{}

func (controller HealthCheckController) RegisterRoutes(router fiber.Router) {
	router.Get("/health", controller.createHealthCheck)
}

func (controller HealthCheckController) createHealthCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(
		dto.GeneralResponse{
			Message: "OK",
		},
	)
}
