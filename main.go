package main

import (
	"github.com/2k4sm/blogPost/internals"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {

	engine := html.New("./ui/html", ".html")

	app := fiber.New(fiber.Config{
		Views:         engine,
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "blogPost",
		AppName:       "blogPost v0.0.1",
	})

	//app.Static("/static", "../../ui/html")

	app.Get("/", internals.ViewAll)
	app.Post("/createpost", internals.CreatePost)
	app.Listen(":8000")

}
