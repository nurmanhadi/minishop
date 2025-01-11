package product

import "context"

type ProductRepository interface {
	InsertProduct(ctx context.Context, product *Product) error
	GetProductById(ctx context.Context, productId string) (*Product, error)
	UpdateProduct(ctx context.Context, product *Product, productId string) error
	DeleteProductById(ctx context.Context, productId string) error
	CountProductById(ctx context.Context, productId string) (int, error)
	GetAllProducts(ctx context.Context) ([]Product, error)
}
type ProductService interface {
	GetAllProducts() ([]Product, error)
	GetProductById(productId string) (*Product, error)
	InsertProduct(body *ProductAddRequestDto) error
}
