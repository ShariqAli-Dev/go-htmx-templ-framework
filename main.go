package main

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/jackc/pgx/v5"
	"github.com/shareed2k/goth_fiber"
	"github.com/shariqali-dev/quizmify/api"
	"github.com/shariqali-dev/quizmify/db"

	"github.com/shariqali-dev/quizmify/views"
)

var (
	baseRoute string
	config    = fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			// 404 error
			if e, ok := err.(*fiber.Error); ok && e.Code == fiber.StatusNotFound {
				extension := filepath.Ext(c.Path())
				// asset placement
				if extension != "" {
					trimmedPath := strings.TrimPrefix(c.Path(), baseRoute)
					correctedPath := "./web/dist/" + trimmedPath
					_, err := os.Stat(c.Path())
					if err != nil {
						return adaptor.HTTPHandler(templ.Handler(views.NotFound()))(c)
					}
					return c.SendFile(correctedPath)
				}
				// 404 page
				lastIndex := strings.LastIndex(c.Path(), "/")
				baseRoute = c.Path()[:lastIndex+1]
				return adaptor.HTTPHandler(templ.Handler(views.NotFound()))(c)
			}

			// generic error
			return adaptor.HTTPHandler(templ.Handler(views.NotFound()))(c)
		},
	}
)

func main() {
	client, err := pgx.Connect(context.Background(), db.DB_URL)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close(context.TODO())

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

	app.Listen(":3000")
}
