package handler

import (
	"errors"
	"os"
	"strconv"

	"gihtub.com/AlifJian/resto-server/database"
	"gihtub.com/AlifJian/resto-server/model"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func AuthLoginUser(c *fiber.Ctx) error {
	requestBody := new(model.UserLoginRequest)
	user := new(model.User)
	response := new(model.UserLoginResponse)

	if errRequestBody := c.BodyParser(requestBody); errRequestBody != nil {
		panic("User handler > " + errRequestBody.Error())
	}

	requestEmailUser := requestBody.Email
	requestPasswordUser := []byte(requestBody.Password)

	result := database.Db.Where("email = ?", requestEmailUser).First(&user)

	if result.Error != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), requestPasswordUser) != nil {
		response = &model.UserLoginResponse{
			Message: "Email or Password didnt Match",
		}
		return c.Status(401).JSON(response)
	}

	response = &model.UserLoginResponse{
		RefreshToken: "Ada",
		AccessToken:  "Ada",
		Message:      "Login Success",
	}

	return c.Status(200).JSON(response)
}

func AuthRegisterUser(c *fiber.Ctx) error {
	BCRYPT_COST := os.Getenv("BCRYPT_COST")
	INT_BCRYPT_COST, errIntBcrypt := strconv.Atoi(BCRYPT_COST)
	response := new(model.UserRegisterResponse)

	if errIntBcrypt != nil {
		panic("User handler > " + errIntBcrypt.Error())
	}

	requestBody := new(model.UserRegisterRequest)

	if errRequestBody := c.BodyParser(requestBody); errRequestBody != nil {
		panic("User handler > " + errRequestBody.Error())
	}

	if requestBody.Password != requestBody.ConfirmPassword {
		response = &model.UserRegisterResponse{
			Message: "Password and Confirm Password didnt match",
		}
		return c.Status(401).JSON(response)
	}

	password := []byte(requestBody.Password)

	hashPassword, errHashPassword := bcrypt.GenerateFromPassword(password, INT_BCRYPT_COST)

	if errHashPassword != nil {
		panic("User handler > " + errHashPassword.Error())
	}

	user := model.User{
		Name:     requestBody.Name,
		Email:    requestBody.Email,
		Password: string(hashPassword),
	}

	result := database.Db.Create(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			response = &model.UserRegisterResponse{
				Message: "Email already registered",
			}
		} else {
			response = &model.UserRegisterResponse{
				Message: "Error Create user",
			}
		}
		return c.Status(500).JSON(response)
	}

	response = &model.UserRegisterResponse{
		User:    user,
		Message: "Success Create User",
	}

	return c.Status(201).JSON(user)
}
