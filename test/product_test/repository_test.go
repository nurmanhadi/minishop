package producttest

import (
	"context"
	"fmt"
	"minishop/external/database/mariadb"
	_ "minishop/internal/config"
	"minishop/internal/features/product"
	"testing"
	"time"
)

var ctx = context.Background()

func TestRepoGetProductById(t *testing.T) {
	db := mariadb.Connection()
	defer db.Close()
	repo := product.NewProductRepository(db)

	product, err := repo.GetProductById(ctx, "1")
	if err != nil {
		panic(err)
	}
	fmt.Println(product)
}
func TestRepoInsertProduct(t *testing.T) {
	db := mariadb.Connection()
	defer db.Close()
	repo := product.NewProductRepository(db)
	product := &product.Product{
		Id:          "2",
		Sku:         "2",
		Name:        "hura hura hura",
		Description: "12345",
		Price:       5000,
		Stock:       1000,
		CreateAt:    time.Now(),
		UpdatedAt:   time.Now(),
	}
	err := repo.InsertProduct(ctx, product)
	if err != nil {
		panic(err)
	}
}
func TestRepoUpdateProduct(t *testing.T) {
	db := mariadb.Connection()
	defer db.Close()
	repo := product.NewProductRepository(db)
	product := &product.Product{
		Name:      "hura",
		UpdatedAt: time.Now(),
	}
	err := repo.UpdateProduct(ctx, product, "2")
	if err != nil {
		panic(err)
	}
}
func TestRepoDeleteProduct(t *testing.T) {
	db := mariadb.Connection()
	defer db.Close()
	repo := product.NewProductRepository(db)

	err := repo.DeleteProductById(ctx, "2")
	if err != nil {
		panic(err)
	}
}
func TestRepoCountProductById(t *testing.T) {
	db := mariadb.Connection()
	defer db.Close()
	repo := product.NewProductRepository(db)
	count, err := repo.CountProductById(ctx, "1")
	if err != nil {
		panic(err)
	}
	fmt.Println(count)
}
func TestRepoGetAllProducts(t *testing.T) {
	db := mariadb.Connection()
	defer db.Close()
	repo := product.NewProductRepository(db)
	products, err := repo.GetAllProducts(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(products)
}
