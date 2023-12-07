package main

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"go.elastic.co/apm/module/apmfiber"
)

func main() {
	app := fiber.New()
	app.Use(apmfiber.Middleware())
	app.Get("/", func(c *fiber.Ctx) error {
		return errors.New("Test Error")
	})
	app.Listen(":8086")
}

