package product

type ProductAddRequestDto struct {
	Sku         string `json:"sku" validate:"required,min=1,max=20"`
	Name        string `json:"name" validate:"required,min=1,max=100"`
	Description string `json:"description" validate:"omitempty,min=1"`
	Price       int    `json:"price" validate:"required"`
	Stock       int    `json:"stock" validate:"required"`
}
type ProductUpdateRequestDto struct {
	Name        string `json:"name" validate:"omitempty,min=1,max=100"`
	Description string `json:"description" validate:"omitempty,min=1"`
	Price       int    `json:"price" validate:"omitempty"`
	Stock       int    `json:"stock" validate:"omitempty"`
}
