package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/shariqali-dev/quizmify/views"
)

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	},
}
var wordMapWords = []string{
	"Hello",
	"world",
	"normally",
	"you",
	"want",
	"more",
	"words",
	"than",
	"this",
	"shariq was here",
}

func main() {
	listenAddr := ":42069"
	jsonWordMapWords, err := json.Marshal(wordMapWords)
	if err != nil {
		log.Fatal(err)
	}
	stringifiedJsonWordMapWords := string(jsonWordMapWords)
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New(config)
	app.Static("/", "./web/dist")
	app.Get("/", adaptor.HTTPHandler(templ.Handler(views.IndexPage())))
	app.Get("/dashboard", adaptor.HTTPHandler(templ.Handler(views.DashboardPage(stringifiedJsonWordMapWords))))
	app.Get("/quiz", adaptor.HTTPHandler(templ.Handler(views.Quiz())))

	fmt.Println(listenAddr)
	app.Listen(listenAddr)
}
