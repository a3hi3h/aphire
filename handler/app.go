package handler

import "github.com/gofiber/fiber/v2"

// Hello hanlde api status
func HomePage(c *fiber.Ctx) error {
	return c.Render("home", fiber.Map{
		"Title": "Hello, World!",
	})
}

func PricingPage(c *fiber.Ctx) error {
	return c.Render("pricing", fiber.Map{
		"Title": "Hello, World!",
	})
}

func AboutPage(c *fiber.Ctx) error {
	return c.Render("about", fiber.Map{
		"Title": "Hello, World!",
	})
}
