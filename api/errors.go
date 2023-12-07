package api

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/shariqali-dev/quizmify/views"
)

var baseRoute string

func ErrorHandler(c *fiber.Ctx, err error) error {
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
}
