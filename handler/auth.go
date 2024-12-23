package handler

import "github.com/gofiber/fiber/v2"

func AuthLoginUser(c *fiber.Ctx) error {

	return c.SendString("Success")
}
