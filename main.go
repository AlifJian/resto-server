package main

import (
	"gihtub.com/AlifJian/resto-server/database"
	"gihtub.com/AlifJian/resto-server/routes"
	"gihtub.com/AlifJian/resto-server/util"
	"github.com/gofiber/fiber/v2"
)

func main() {
	defer util.Catch()
	database.InitDb()

	app := fiber.New()

	routes.InitAuthRoute(app)

	app.Listen(":3080")
}
