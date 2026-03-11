// golang gin framework mvc and clean code project
// Licensed under the Apache License 2.0
// @author Selman TUNÇ <selmantunc@gmail.com>
// @link: https://github.com/stnc/go-mvc-blog-clean-code
// @license: Apache License 2.0
package main

import (
	
	"github.com/joho/godotenv"
	
	"log"



	"github.com/gofiber/fiber/v3"

	"github.com/gofiber/template/django/v4"

	"helix/app/handlers"
)

func init() {
	//To load our environmental variables.

	if err := godotenv.Load(); err != nil {
		log.Println("no env gotten")
	}
}

func main() {
// Middleware
	engine := django.New("./views", ".html")
	engine.Reload(true)
	// Pass the engine to the Views
	app := fiber.New(fiber.Config{Views: engine})



	handlers.Setup(app)

	log.Fatal(app.Listen(":9999"))

	

}
