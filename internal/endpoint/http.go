package endpoint

import (
	"fmt"
	"sync"
	"github.com/gofiber/fiber/v2"
	cont "github.com/ghabxph/xv94dx3/internal/controller"
)

type Http struct {
	Port int
}

// This function is blocking.
func (h *Http) Create(wg *sync.WaitGroup) {
	defer wg.Done()

	app := fiber.New()

	// Routes
	app.Get("/top/confirmed", getTopConfirmed(&cont.TopConfirmed{}))

	// Run http server
	app.Listen(fmt.Sprintf(":%d", h.Port))
}

func getTopConfirmed(controller *cont.TopConfirmed) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return c.JSON(controller.GetTopConfirmed("1", 1))
	}
}
