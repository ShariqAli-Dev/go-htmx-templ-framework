package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shariqali-dev/quizmify/db"
)

type PersonHandler struct {
	personStore db.PersonStore
}

func NewPersonHandler(personStore db.PersonStore) *PersonHandler {
	return &PersonHandler{
		personStore: personStore,
	}
}

func (h *PersonHandler) HandleGetPeople(c *fiber.Ctx) error {
	people, err := h.personStore.GetPeople()
	if err != nil {
		return err
	}
	return c.JSON(people)
}
