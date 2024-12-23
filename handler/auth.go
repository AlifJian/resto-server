package handler

import (
	"gihtub.com/AlifJian/resto-server/database"
	"gihtub.com/AlifJian/resto-server/model"
	"github.com/gofiber/fiber/v2"
)

func AuthLoginUser(c *fiber.Ctx) error {
	requestBody := new(model.UserLoginRequest)
	user := new(model.User)
	response := new(model.UserLoginResponse)

	if errRequestBody := c.BodyParser(requestBody); errRequestBody != nil {
		panic("User handler > " + errRequestBody.Error())
	}

	emailUser := requestBody.Email
	result := database.Db.Where("email = ?", emailUser).First(&user)

	if result.Error != nil || user.Password != requestBody.Password {
		return c.Status(401).JSON(response)
	}

	response = &model.UserLoginResponse{
		RefreshToken: "Ada",
		AccessToken:  "Ada",
	}

	return c.Status(200).JSON(response)
}
