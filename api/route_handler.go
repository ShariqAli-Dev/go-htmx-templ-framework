package api

import (
	"encoding/json"
	"log"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/shariqali-dev/quizmify/views"
)

type RouteHandler struct {
}

func NewRouteHandler() *RouteHandler {
	return &RouteHandler{}
}
func (h *RouteHandler) HandleGetIndex(c *fiber.Ctx) error {
	return adaptor.HTTPHandler(templ.Handler(views.IndexPage()))(c)
}

func (h *RouteHandler) HandleGetDashboard(c *fiber.Ctx) error {
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
	jsonWordMapWords, err := json.Marshal(wordMapWords)
	if err != nil {
		log.Fatal(err)
	}
	stringifiedJsonWordMapWords := string(jsonWordMapWords)
	if err != nil {
		log.Fatal(err)
	}
	return adaptor.HTTPHandler(templ.Handler(views.DashboardPage(stringifiedJsonWordMapWords)))(c)
}

func (h *RouteHandler) HandleGetQuiz(c *fiber.Ctx) error {
	return adaptor.HTTPHandler(templ.Handler(views.Quiz()))(c)
}
