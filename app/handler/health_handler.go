package handler

import (
	"btc/app/service"
	"github.com/gofiber/fiber/v2"
)

func NewHealthHandler(app fiber.Router, healthSrv service.IHealthService) {
	app.Get("/health", HealthCheck(healthSrv))
}

// HealthCheck func get health status or 500.
// @Description Get health status.
// @Summary get health status or 500
// @Tags Health
// @Accept json
// @Produce json
// @Success 200
// @Router /v1/health [get]
func HealthCheck(healthService service.IHealthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := healthService.HealthCheck()

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}

		// Return status 200 OK.
		return c.JSON(fiber.Map{
			"error": false,
			"msg":   nil,
		})
	}
}
