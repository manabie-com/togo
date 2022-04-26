package utils

import "github.com/gofiber/fiber/v2"

//BodyParser ...
func BodyParser(c *fiber.Ctx, in interface{}) error {

	err := c.BodyParser(in)
	if err != nil {
		c.Status(500).JSON(fiber.Map{"retCode": 500, "message": err.Error()})
	}
	return err
}
