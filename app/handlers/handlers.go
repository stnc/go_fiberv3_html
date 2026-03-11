package handlers

import (
    "github.com/gofiber/fiber/v3"
)

// Home renders the home view
// Home renders the home view
func Home(c fiber.Ctx) error {
	return c.Render("homepage", fiber.Map{
		"Title": "Hello, World!",
	})
}


func Users(c fiber.Ctx) error {
	return c.Render("users", fiber.Map{
		"Title": "Hello, usersss !",
	})
}


// About renders the about view
func About(c fiber.Ctx) error {
    return c.Render("about", nil)  // Boş data için nil daha temiz
}

// NotFound renders the 404 view
func NotFound (c fiber.Ctx) error {
    return c.Status(fiber.StatusNotFound).Render("404", nil)
}