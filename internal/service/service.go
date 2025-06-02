package service

import (
	"nexa/internal/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func StartServer(uh *handler.UserHandler) {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	api := app.Group("/api/users")

	api.Post("/", uh.CreateUser)
	api.Post("/login", uh.LoginUser)
	api.Get("/", uh.GetUsers)
	api.Get("/:id", uh.GetUserByID)
	api.Put("/:id", uh.UpdateUser)
	api.Delete("/:id", uh.DeleteUser)

	app.Listen(":8080")
}
