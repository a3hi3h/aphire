package main

import (
	"churnsight/database"
	"churnsight/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	swagger "github.com/gofiber/swagger"
	"github.com/gofiber/template/html"
)

// @title Aperno Hiring API
// @version 1.0
// @description This is a Aperno API
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	engine := html.New("./templates", ".tmpl")
	engine.Reload(true)
	//engine.Debug(true)
	app := fiber.New(fiber.Config{
		Views:             engine,
		EnablePrintRoutes: false,
	})

	app.Static("/assets", "templates/assets")
	app.Get("/swagger/*", swagger.HandlerDefault) // default
	app.Use(cors.New())
	app.Use(logger.New())
	app.Get("/dashboard", monitor.New())

	database.ConnectDB()

	router.SetupRoutes(app)
	log.Fatal(app.Listen(":80"))
}
