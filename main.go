package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"github.com/shareed2k/goth_fiber"
	"github.com/shariqali-dev/quizmify/api"
	"github.com/shariqali-dev/quizmify/db"
	"github.com/shariqali-dev/quizmify/types"
)

func main() {
	client, err := gorm.Open("postgres", db.DB_URL)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	client.AutoMigrate(&types.Person{})
	client.AutoMigrate(&types.Book{})
	if err := client.Create(&types.Person{
		Name:  "shariq",
		Email: "hehe",
		Books: []types.Book{
			{
				Title:      "cool",
				Author:     "shariq",
				CallNumber: 234,
				PersonID:   1,
			},
		},
	}).Error; err != nil {
		log.Fatal(err)
	}

	var (
		personStore = db.NewPersonStore(client)

		store = &db.Store{
			PersonStore: *personStore,
		}

		routeHandler  = api.NewRouteHandler()
		authHandler   = api.NewAuthHandler()
		personHandler = api.NewPersonHandler(&store.PersonStore)

		app = fiber.New(fiber.Config{
			ErrorHandler: api.ErrorHandler,
		})
		apiv1 = app.Group("/api")
	)

	app.Static("/", "./web/dist")
	// routing
	app.Get("/login/:provider", goth_fiber.BeginAuthHandler)
	app.Get("/auth/:provider/callback", authHandler.HandleGetAuthorizationCallback)
	app.Get("/logout", authHandler.HandleGetLogout)
	app.Get("/", routeHandler.HandleGetIndex)
	app.Get("/dashboard", api.ValidateUser, routeHandler.HandleGetDashboard)
	app.Get("/quiz", api.ValidateUser, routeHandler.HandleGetQuiz)

	// api
	apiv1.Get("/people", personHandler.HandleGetPeople)

	app.Listen(":3000")
}
