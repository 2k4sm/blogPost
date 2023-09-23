package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {

	engine := html.New("../../ui/html", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	//app.Static("/static", "../../ui/html")

	app.Get("/", homeHandler)

	app.Get("/api/viewall", viewAll)

	app.Listen(":8000")

}
