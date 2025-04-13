package healthcheckflow

import (
	config "backend_project/configs"
	"backend_project/internal/middleware/exception"
	"backend_project/internal/models/entity"

	"github.com/gofiber/fiber/v2"
)

func GetHealthCheck(c *fiber.Ctx) error {
	var healthCheck []entity.HealthCheckTable
	if err := config.DB.Find(&healthCheck).Error; err != nil {
		exception.PanicLog(err)
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(fiber.Map{"status": healthCheck[0].Status})
}
