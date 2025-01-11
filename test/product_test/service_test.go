package producttest

import (
	"fmt"
	"minishop/external/database/mariadb"
	_ "minishop/internal/config"
	"minishop/internal/features/product"
	"testing"
	"time"

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

	product, err := serv.GetProductById("1")
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
		// Price:       12000,
		// Stock:       100,
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
	product := &product.Product{
		Name:      "hura",
		UpdatedAt: time.Now(),
	}
	err := repo.UpdateProduct(ctx, product, "2")
	if err != nil {
		panic(err)
	}
}
func TestServDeleteProduct(t *testing.T) {
	db := mariadb.Connection()
	defer db.Close()
	repo := product.NewProductRepository(db)

	err := repo.DeleteProductById(ctx, "2")
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
