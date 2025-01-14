package di

import (
	"context"
	"minishop/external/database/mariadb"
	"minishop/internal/features/product"

	_ "minishop/internal/config"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func DependencyInjection(app *fiber.App) {
	valodation := validator.New()
	ctx := context.Background()

	db := mariadb.Connection()

	productRepository := product.NewProductRepository(db)
	productService := product.NewProductService(&productRepository, ctx, valodation)
	productHandler := product.NewProductHandler(&productService)
	product.ProductRoute(app, productHandler)
}
