package exceptions

import "github.com/gofiber/fiber/v2"

type GlobalErrorHandlerResp struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func ErorrHandler(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusBadRequest).JSON(GlobalErrorHandlerResp{
		Success: false,
		Message: err.Error(),
	})
}
