package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shareed2k/goth_fiber"
	"github.com/shariqali-dev/quizmify/api"
)

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {
	listenAddr := ":3000"

	var (
		routeHandler = api.NewRouteHandler()
		authHandler  = api.NewAuthHandler()
		app          = fiber.New(config)
	)

	app.Static("/", "./web/dist")
	app.Get("/login/:provider", goth_fiber.BeginAuthHandler)
	app.Get("/auth/:provider/callback", authHandler.HandleGetAuthorizationCallback)
	app.Get("/logout", authHandler.HandleGetLogout)

	app.Get("/", routeHandler.HandleGetIndex)
	app.Get("/dashboard", routeHandler.HandleGetDashboard)
	app.Get("/quiz", routeHandler.HandleGetQuiz)

	app.Listen(listenAddr)
}
