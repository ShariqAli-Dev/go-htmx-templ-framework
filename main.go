package main

import (
	"fmt"

	"github.com/a-h/templ"
	"github.com/a-h/templ-examples/hello-world/views"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {
	listenAddr := ":3000"
	var (
		app = fiber.New(config)
	)
	app.Static("/", "./src/public")
	app.Get("/", adaptor.HTTPHandler(templ.Handler(views.IndexPage())))
	fmt.Println(listenAddr)
	app.Listen(listenAddr)
}
