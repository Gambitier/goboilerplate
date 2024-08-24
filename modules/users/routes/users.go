package routes

import (
	"github.com/gambitier/gocomm/appContext"
	"github.com/gambitier/gocomm/middlewares"
	"github.com/gambitier/gocomm/modules/users/dto"
	"github.com/gambitier/gocomm/modules/users/handlers"
	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(app *appContext.AppContext, httpServer *fiber.App) {
	userHandler := handlers.NewUserHandler(app.DbRepo, app.TempFileStorage, app.MessageQueue)
	users := httpServer.Group("/users")
	users.Post("/", middlewares.ValidateRequest(dto.CreateUserRequest{}), userHandler.CreateUser)
	users.Post("/avatar", userHandler.UploadAvatar)
	users.Get("/profile", GetProfile)
}

func GetProfile(c *fiber.Ctx) error {
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Profile details",
		"data":    "Profile does not exist",
	})
}
