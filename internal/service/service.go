package service

import (
	"nexa/internal/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jackc/pgx/v5"
)

func StartServer(conn *pgx.Conn) {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	uh := handler.NewUserHandler(conn)
	api := app.Group("/api/users")

	api.Post("/", uh.CreateUser)
	api.Get("/", uh.GetUsers)
	api.Get("/:id", uh.GetUserByID)
	api.Put("/:id", uh.UpdateUser)
	// api.Delete("/:id", handler.DeleteUser)

	app.Listen(":8080")
}
