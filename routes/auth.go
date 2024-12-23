package routes

import (
	"gihtub.com/AlifJian/resto-server/handler"
	"github.com/gofiber/fiber/v2"
)

func InitAuthRoute(app *fiber.App) {
	apiAuthGroup := app.Group("/auth")
	apiAuthGroup.Get("/", handler.AuthLoginUser)
}
