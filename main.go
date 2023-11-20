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
	app.Static("/assets", "./src/public/assets")
	app.Get("/", adaptor.HTTPHandler(templ.Handler(views.IndexPage())))
	app.Get("/dashboard", adaptor.HTTPHandler(templ.Handler(views.DashboardPage())))

	fmt.Println(listenAddr)
	app.Listen(listenAddr)
}
