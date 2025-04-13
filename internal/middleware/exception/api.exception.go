package exception

import (
	"backend_project/internal/models/exception"
	"log"

	"github.com/gofiber/fiber/v2"
)

func ApiException(c *fiber.Ctx, err error) error {

	statusCode := fiber.StatusInternalServerError
	errorResponse := exception.ApiError{
		ErrorResponse: exception.ErrorResponse{
			Code:    statusCode,
			Message: "Internal Server Error",
			Details: nil,
		},
	}

	if apiError, ok := err.(*exception.ApiError); ok {
		statusCode = apiError.ErrorResponse.Code
		errorResponse = *apiError
	} else if e, ok := err.(*fiber.Error); ok {
		statusCode = e.Code
		errorResponse.ErrorResponse.Code = statusCode
		errorResponse.ErrorResponse.Message = e.Message
	}

	log.Printf("\033[31mError: %v\033[0m", err)

	return c.Status(statusCode).JSON(errorResponse)
}
