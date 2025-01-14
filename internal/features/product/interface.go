package product

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

type ProductRepository interface {
	InsertProduct(ctx context.Context, product *Product) error
	GetProductById(ctx context.Context, productId *string) (*Product, error)
	UpdateProduct(ctx context.Context, product *Product, productId *string) error
	DeleteProductById(ctx context.Context, productId *string) error
	CountProductById(ctx context.Context, productId *string) (int, error)
	GetAllProducts(ctx context.Context) ([]Product, error)
}
type ProductService interface {
	GetAllProducts() ([]Product, error)
	GetProductById(productId *string) (*Product, error)
	InsertProduct(body *ProductAddRequestDto) error
	UpdateProduct(productId *string, body *ProductUpdateRequestDto) error
	DeleteProductById(productId *string) error
}
type ProductHandler interface {
	GetAllProducts(c *fiber.Ctx) error
	GetProductById(c *fiber.Ctx) error
	InsertProduct(c *fiber.Ctx) error
	UpdateProduct(c *fiber.Ctx) error
	DeleteProduct(c *fiber.Ctx) error
}
