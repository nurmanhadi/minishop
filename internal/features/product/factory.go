package product

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

func NewProductRepository(db *sql.DB) ProductRepository {
	return &repositoryImpl{db: db}
}
func NewProductService(repo *ProductRepository, ctx context.Context, validation *validator.Validate) ProductService {
	return &serviceImpl{repo: *repo, ctx: ctx, validation: validation}
}
func NewProductHandler(serv *ProductService) ProductHandler {
	return &handlerImpl{serv: *serv}
}
