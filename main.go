package main

import (
	"gihtub.com/AlifJian/resto-server/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.InitAuthRoute(app)

	app.Listen(":3080")
}
