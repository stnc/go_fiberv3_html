package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func Setup(app *fiber.App) {
	app.Use(cors.New())
	app.Use(recover.New())
	app.Use(logger.New())

	app.Get("/", func(c fiber.Ctx) error {
		// Render with and extends
		return c.Render("homepage", fiber.Map{
			"Title": "this homepage",
		})
	})	
	
	app.Get("/online2", func(c fiber.Ctx) error {
		// Render with and extends
		return c.Render("online-managemen", fiber.Map{
			"Title": "this homepage",
		})
	})

	app.Get("/embed", func(c fiber.Ctx) error {
		// Render index within layouts/main
		return c.Render("embed", fiber.Map{
			"Title": "Hello, World!",
		}, "layouts/main2")
	})

	app.Get("/public*", static.New("./public"))

	app.Get("/home", Home)
	app.Get("users", Users)

	app.Get("/about", About)


	v1 := app.Group("/v1")
	v1.Get("/users", Users)     // /api/v1/users
	v1.Post("/users", Users)  // /api/v1/users




	app.Use(NotFound) // 404 

}
