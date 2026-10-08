package main

import (
	"github.com/gofiber/fiber/v3"
)

func main() {
	// สร้าง instance ของ Fiber
	app := fiber.New()

	// กําหนด route และ handler
	app.Get("/hellogo", func(c fiber.Ctx) error {
		return c.SendString("Hello Go ^_^ ....")
	})

	app.Get("/higo", func(c fiber.Ctx) error {
		return c.SendString("Hi Go ^_^ ....")
	})

	// start sever
	if err := app.Listen(":4040"); err != nil {
		panic(err)
	}
}
