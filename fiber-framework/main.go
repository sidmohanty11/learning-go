package main

import (
	"github.com/gofiber/fiber"
)

func helloWorld(c *fiber.Ctx) {
	c.Send("hello, world")
}

func main() {
	app := fiber.New()

	//setupRoutes(app)

	app.Listen(3000)
}
