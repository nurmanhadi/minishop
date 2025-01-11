package product

import "time"

type Product struct {
	Id          string
	Sku         string
	Name        string
	Description string
	Price       int
	Stock       int
	CreateAt    time.Time
	UpdatedAt   time.Time
}
