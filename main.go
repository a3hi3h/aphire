package main

import (
	"api-fiber-gorm/database"
	"api-fiber-gorm/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/template/html"
)

func main() {
	engine := html.New("./templates", ".tmpl")
	//engine.Reload(true)
	//engine.Debug(true)
	app := fiber.New(fiber.Config{
		Views:             engine,
		EnablePrintRoutes: false,
	})

	app.Static("/assets", "templates/assets")

	app.Use(cors.New())
	app.Use(logger.New())
	app.Get("/dashboard", monitor.New())

	database.ConnectDB()

	router.SetupRoutes(app)
	log.Fatal(app.Listen(":8080"))
}
