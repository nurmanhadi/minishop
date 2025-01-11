package product

import (
	"context"
	"database/sql"
	"fmt"
	"minishop/pkg/infrastructure/query"
)

type repositoryImpl struct {
	db *sql.DB
}

func (r *repositoryImpl) GetAllProducts(ctx context.Context) ([]Product, error) {
	var products []Product
	stmt, err := r.db.PrepareContext(ctx, query.GetAllProducts)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var product Product
		err := rows.Scan(&product.Id, &product.Sku, &product.Name, &product.Description, &product.Price, &product.Stock, &product.CreateAt, &product.UpdatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
func (r *repositoryImpl) GetProductById(ctx context.Context, productId string) (*Product, error) {
	product := new(Product)
	stmt, err := r.db.PrepareContext(ctx, query.GetProductById)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.QueryContext(ctx, productId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&product.Id, &product.Sku, &product.Name, &product.Description, &product.Price, &product.Stock, &product.CreateAt, &product.UpdatedAt)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("product %s not found", productId)
	}
	return product, nil
}
func (r *repositoryImpl) InsertProduct(ctx context.Context, product *Product) error {
	stmt, err := r.db.PrepareContext(ctx, query.InsertProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, product.Id, product.Sku, product.Name, product.Description, product.Price, product.Stock, product.CreateAt, product.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}
func (r *repositoryImpl) UpdateProduct(ctx context.Context, product *Product, productId string) error {
	args := []interface{}{}
	if product.Name != "" {
		query.UpdateProduct += "name = ?, "
		args = append(args, product.Name)
	}
	if product.Description != "" {
		query.UpdateProduct += "description = ?, "
		args = append(args, product.Description)
	}
	if product.Price != 0 {
		query.UpdateProduct += "price = ?, "
		args = append(args, product.Price)
	}
	if product.Stock != 0 {
		query.UpdateProduct += "stock = ?, "
		args = append(args, product.Stock)
	}
	if !product.UpdatedAt.IsZero() {
		query.UpdateProduct += "updated_at = ?, "
		args = append(args, product.UpdatedAt)
	}
	query.UpdateProduct = query.UpdateProduct[:len(query.UpdateProduct)-2] + " WHERE id = ?"
	args = append(args, productId)
	stmt, err := r.db.PrepareContext(ctx, query.UpdateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, args...)
	if err != nil {
		return err
	}
	return nil
}
func (r *repositoryImpl) DeleteProductById(ctx context.Context, productId string) error {
	stmt, err := r.db.PrepareContext(ctx, query.DeleteProductById)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, productId)
	if err != nil {
		return err
	}
	return nil
}
func (r *repositoryImpl) CountProductById(ctx context.Context, productId string) (int, error) {
	var count int
	stmt, err := r.db.PrepareContext(ctx, query.CountProductById)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	rows, err := stmt.QueryContext(ctx, productId)
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			return 0, err
		}
	}

	return count, nil
}
