package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("/api/viewall", viewAll)

	app.Listen(":8000")

}
