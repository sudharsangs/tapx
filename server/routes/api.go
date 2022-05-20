package routes

import (
	"github.com/gofiber/fiber/v2"
	Controller "github.com/sudharsangs/tapx/app/controllers/api"
)

func SearchApi(api fiber.Router) {
	searchApi := api.Group("/")
	searchApi.Get("/results", Controller.GetScrapingResults)
}
