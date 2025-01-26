package utils

import (
	"github.com/gofiber/fiber/v2"
)

type ApiResponse struct {
	Status  bool        `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

func SendApiResponse(c *fiber.Ctx, status bool, code int, message string, data interface{}, errors interface{}) error {

	response := ApiResponse{
		Status:  status,
		Code:    code,
		Message: message,
		Data:    data,
		Errors:  errors,
	}

	return c.Status(code).JSON(response)
}
