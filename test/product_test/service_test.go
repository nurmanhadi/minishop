package producttest

import (
	"fmt"
	"minishop/external/database/mariadb"
	_ "minishop/internal/config"
	"minishop/internal/features/product"
	"testing"

	"github.com/go-playground/validator/v10"
)

var validation = validator.New()

func TestServGetAllProducts(t *testing.T) {
	db := mariadb.Connection()
	defer db.Close()
	repo := product.NewProductRepository(db)
	serv := product.NewProductService(&repo, ctx, validation)

	products, err := serv.GetAllProducts()
	if err != nil {
		panic(err)
	}
	fmt.Println(products)
}
func TestServGetProductById(t *testing.T) {
	db := mariadb.Connection()
	defer db.Close()
	repo := product.NewProductRepository(db)
	serv := product.NewProductService(&repo, ctx, validation)

	product, err := serv.GetProductById("2")
	if err != nil {
		panic(err)
	}
	fmt.Println(product)
}
func TestServInsertProduct(t *testing.T) {
	db := mariadb.Connection()
	defer db.Close()
	repo := product.NewProductRepository(db)
	serv := product.NewProductService(&repo, ctx, validation)
	body := &product.ProductAddRequestDto{
		Sku:         "",
		Name:        "",
		Description: "",
		Price:       12000,
		Stock:       100,
	}
	err := serv.InsertProduct(body)
	if err != nil {
		panic(err)
	}
}
func TestServUpdateProduct(t *testing.T) {
	db := mariadb.Connection()
	defer db.Close()
	repo := product.NewProductRepository(db)
	serv := product.NewProductService(&repo, ctx, validation)
	product := &product.ProductUpdateRequestDto{
		Name:        "yarezo",
		Description: "",
		Price:       0,
		Stock:       0,
	}
	err := serv.UpdateProduct("2", product)
	if err != nil {
		panic(err)
	}
}
func TestServDeleteProduct(t *testing.T) {
	db := mariadb.Connection()
	defer db.Close()
	repo := product.NewProductRepository(db)
	serv := product.NewProductService(&repo, ctx, validation)
	err := serv.DeleteProductById("2")
	if err != nil {
		panic(err)
	}
}
func TestServCountProductById(t *testing.T) {
	db := mariadb.Connection()
	defer db.Close()
	repo := product.NewProductRepository(db)
	count, err := repo.CountProductById(ctx, "1")
	if err != nil {
		panic(err)
	}
	fmt.Println(count)
}
