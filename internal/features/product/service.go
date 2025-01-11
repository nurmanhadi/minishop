package product

import (
	"context"
	"minishop/pkg/infrastructure/exception"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type serviceImpl struct {
	repo       ProductRepository
	ctx        context.Context
	validation *validator.Validate
}

func (s *serviceImpl) GetAllProducts() ([]Product, error) {
	products, err := s.repo.GetAllProducts(s.ctx)
	if err != nil {
		return nil, err
	}
	return products, nil
}
func (s *serviceImpl) GetProductById(productId string) (*Product, error) {
	product, err := s.repo.GetProductById(s.ctx, productId)
	if err != nil {
		return nil, exception.ProductNotFound
	}
	return product, nil
}
func (s *serviceImpl) InsertProduct(body *ProductAddRequestDto) error {
	if err := s.validation.Struct(body); err != nil {
		exception.ValidationError(err)
		return exception.ProductValidationError
	}
	id := uuid.New().String()
	product := &Product{
		Id:          id,
		Sku:         body.Sku,
		Name:        body.Name,
		Description: body.Description,
		Price:       body.Price,
		Stock:       body.Stock,
		CreateAt:    time.Now(),
		UpdatedAt:   time.Now(),
	}
	err := s.repo.InsertProduct(s.ctx, product)
	if err != nil {
		return err
	}
	return nil
}
