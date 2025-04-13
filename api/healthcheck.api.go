package api

import (
	"backend_project/internal/flow/healthcheckflow"

	"github.com/gofiber/fiber/v2"
)

func InitHealthCheckRoute(app *fiber.App) {
	app.Get("/healthcheck", healthcheckflow.GetHealthCheck)
	app.Get("/error", healthcheckflow.GetApiException)
}
