package query

var (
	GetAllProducts    = "SELECT id, sku, name, description, price, stock, create_at, updated_at FROM products"
	InsertProduct     = "INSERT INTO products(id, sku, name, description, price, stock, create_at, updated_at) VALUES(?,?,?,?,?,?,?,?)"
	UpdateProduct     = "UPDATE products SET "
	GetProductById    = "SELECT id, sku, name, description, price, stock, create_at, updated_at FROM products WHERE id = ? LIMIT 1"
	DeleteProductById = "DELETE FROM products WHERE id = ? LIMIT 1"
	CountProductById  = "SELECT COUNT(*) FROM products WHERE id = ? LIMIT 1"
)
