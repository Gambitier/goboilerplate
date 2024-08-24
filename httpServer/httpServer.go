package httpserver

import (
	"github.com/gambitier/gocomm/appContext"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type HttpServer struct {
	App        *fiber.App
	AppContext *appContext.AppContext
}

func NewHttpServer(appContext *appContext.AppContext) *HttpServer {
	return &HttpServer{
		App:        fiber.New(),
		AppContext: appContext,
	}
}

func (server *HttpServer) Configure() {
	server.App.Use(etag.New())
	server.App.Use(cache.New())
	server.App.Use(compress.New())
	server.App.Use(recover.New())
	server.App.Use(idempotency.New())
	server.App.Use(func(c *fiber.Ctx) error {
		if c.Path() == "/swagger" {
			return c.Next()
		}
		return helmet.New()(c)
	})

	// Middleware
	server.App.Use(logger.New())

	// Config swagger
	server.App.Use(swagger.New(swagger.Config{
		BasePath: "/",
		FilePath: "./_apidocs/swagger.json",
		Path:     "swagger",
		Title:    "Swagger API Docs",
	}))
}
