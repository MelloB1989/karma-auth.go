package routes

import (
	"github.com/gofiber/fiber/v2"
	user "karma_auth/handlers/users"
	middlewares "karma_auth/middlewares"
)

func Users() *fiber.App {
	app := fiber.New()
	v1 := app.Group("/v1")

	//User routes
	users := v1.Group("/users")
	users.Get("/h", user.GetUsers)
	users.Post("/create", middlewares.IsKarmaOrganisation, user.CreateUser)

	//Organisation routes
	organisations := v1.Group("/organisations")
	organisations.Get("/create", user.CreateUser)

	return app
}