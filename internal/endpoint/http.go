package endpoint

import (
	"fmt"
	"sync"
	"strconv"
	"github.com/gofiber/fiber/v2"
	"github.com/arsmn/fiber-swagger/v2"
	cont "github.com/ghabxph/xv94dx3/internal/controller"
)

type Http struct {
	Port int
}

// This function is blocking.
func (h *Http) Create(wg *sync.WaitGroup) {
	defer wg.Done()

	app := fiber.New()

	// Swagger
	app.Get("/", func (c *fiber.Ctx) error { return c.Redirect("/docs/index.html") })
	app.Get("/docs/*", swagger.New(swagger.Config{DeepLinking: true, URL:"/swagger.yaml"}))
	app.Static("/swagger.yaml", "./swagger.yaml")

	// Routes
	app.Get("/top/confirmed", getTopConfirmed(&cont.TopConfirmed{}))

	// Run http server
	app.Listen(fmt.Sprintf(":%d", h.Port))
}

func getTopConfirmed(controller *cont.TopConfirmed) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		observationDate := c.Query("observation_date", "")
		_maxResults := c.Query("max_results", "10")

		// observation_date is required
		if observationDate == "" {
			return c.JSON(map[string]string{
				"error": "observation_date parameter is required.",
				"usage": "url/top/confirmed?observation_date=yyyy-mm-dd",
			})
		}


		maxResults, err := strconv.ParseUint(_maxResults, 10, 32)

		// max_results must be integer
		if err != nil {
			return c.JSON(map[string]string{"error": "max_results is not integer."})
		}

		// max_results must be positive integer.
		if maxResults < 0 {
			return c.JSON(map[string]string{"error": "max_results must be positive integer."})
		}
		return c.JSON(controller.GetTopConfirmed(observationDate, uint(maxResults)))
	}
}
