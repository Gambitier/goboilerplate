package errors

import "github.com/gofiber/fiber/v2"

var (
	UserNotFound       = fiber.NewError(fiber.StatusNotFound, "user not found")
	FailedToCreateUser = fiber.NewError(fiber.StatusNotFound, "failed to create user")
)
