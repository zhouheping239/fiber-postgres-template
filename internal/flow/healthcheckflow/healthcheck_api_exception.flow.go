package healthcheckflow

import (
	"backend_project/internal/models/exception"

	"github.com/gofiber/fiber/v2"
)

func GetApiException(c *fiber.Ctx) error {
	details := []exception.ErrorDetail{
		{
			Reason:    "reason",
			Reference: []string{"/error"},
		},
	}
	return &exception.ApiError{
		ErrorResponse: exception.ErrorResponse{
			Code:    fiber.StatusNotFound,
			Message: "example error",
			Details: &details,
		},
	}
}
