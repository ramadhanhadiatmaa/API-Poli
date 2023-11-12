package routes

import (
	"apipoli/controllers"
	"apipoli/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Route(r *fiber.App) {

	kamar := r.Group("/api")

	kamar.Get("/", middlewares.Auth, controllers.Index)
	kamar.Get("/:kode_poli", middlewares.Auth, controllers.Show)
	kamar.Post("/", middlewares.Auth, controllers.Create)
	kamar.Put("/:kode_poli", middlewares.Auth, controllers.Update)
	kamar.Delete("/:kode_poli", middlewares.Auth, controllers.Delete)
}