package product

import "github.com/gofiber/fiber/v2"

func ProductRoute(app *fiber.App, h ProductHandler) {
	product := app.Group("products")
	product.Get("/", h.GetAllProducts)
	product.Get("/:productId", h.GetProductById)
	product.Post("/", h.InsertProduct)
	product.Patch("/:productId", h.UpdateProduct)
	product.Delete("/:productId", h.DeleteProduct)
}
