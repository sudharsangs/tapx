package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sudharsangs/tapx/routes"
)

func main() {
	app := fiber.New()
	api := app.Group("api")
	apiv1 := api.Group("/v1")
	routes.SearchApi(apiv1)

	var err error

	app.Use(func(c *fiber.Ctx) error {
		if err := c.SendStatus(fiber.StatusNotFound); err != nil {
			panic(err)
		}
		if err := c.Render("errors/404", fiber.Map{}); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return err
	})

	// Start listening on the specified address
	err = app.Listen(":5600")
}
