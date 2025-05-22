package utils

import (
	"github.com/gofiber/fiber/v2"
)

type SuccessStructure struct {
	StatusCode int
	Message    string
	Data       interface{}
}

type ErrorStructure struct {
	StatusCode int
	Message    string
	Error      interface{}
}

func HttpSuccess(c *fiber.Ctx, s SuccessStructure) error {
	return c.Status(s.StatusCode).JSON(fiber.Map{
		"message": s.Message,
		"data":    s.Data,
	})

}

func HttpError(c *fiber.Ctx, e ErrorStructure) error {
	return c.Status(e.StatusCode).JSON(fiber.Map{
		"message": e.Message,
		"error":   e.Error,
	})
}
