package validation

import "github.com/gofiber/fiber/v2"

type ValidationError struct {
	FailedField    string `json:"failedField"`
	ValidatorKey   string `json:"validatorKey"`
	ValidatorParam string `json:"validatorParam"`
}

func FailedValidationResponse(c *fiber.Ctx, validationErrors []*ValidationError) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"type":  "validation-error",
		"items": validationErrors,
	})
}
