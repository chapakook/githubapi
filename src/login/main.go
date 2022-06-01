package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	engine := html.New("./pages", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/public", "./public")

	app.Get("/", Index)
	app.Get("/oauth", Oauth)
	app.Get("/end", End)

	log.Fatal(app.Listen(PORT))
}
