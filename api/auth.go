package api

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/sqlite3"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"github.com/shareed2k/goth_fiber"
)

type AuthHandler struct {
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) HandleGetAuthorizationCallback(c *fiber.Ctx) error {
	user, err := goth_fiber.CompleteUserAuth(c, goth_fiber.CompleteUserAuthOptions{ShouldLogout: false})
	if err != nil {
		log.Fatal(err)
		return c.JSON(err)
	}

	fmt.Println(user)
	return c.Redirect("/dashboard", fiber.StatusFound)

}

func (h *AuthHandler) HandleGetLogout(c *fiber.Ctx) error {
	if err := goth_fiber.Logout(c); err != nil {
		log.Fatal(err)
	}

	return c.SendString("logout")
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env files")
	}

	googleClientId := os.Getenv("GOOGLE_CLIENT_ID")
	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")

	config := session.Config{
		Storage:      sqlite3.New(),
		KeyLookup:    "cookie:_gothic_session",
		CookieSecure: os.Getenv("ENVIRONMENT") == "production",
	}
	sessions := session.New(config)
	goth_fiber.SessionStore = sessions

	goth.UseProviders(
		google.New(googleClientId, googleClientSecret, "http://localhost:3000/auth/google/callback"))

}
