package main

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/shareed2k/goth_fiber"
	"github.com/shariqali-dev/quizmify/api"
	"github.com/shariqali-dev/quizmify/views"
)

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return adaptor.HTTPHandler(templ.Handler(views.NotFound()))(c)
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
	app.Get("/dashboard", api.ValidateUser, routeHandler.HandleGetDashboard)
	app.Get("/quiz", api.ValidateUser, routeHandler.HandleGetQuiz)

	app.Listen(listenAddr)
}
