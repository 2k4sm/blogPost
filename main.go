package main

import (
	"fmt"
	"os"

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

	app.Static("/", "./ui/assets")

	app.Get("/", internals.RenderHome)
	key := os.Getenv("API_KEY")
	app.Get(fmt.Sprintf("/api/%s/views", key), internals.ViewAll)
	app.Get("/createpost", internals.CreatePost)
	app.Post("/createpost", internals.ProcessForm)
	app.Listen(":8000")

}
