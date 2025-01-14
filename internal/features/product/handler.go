package product

import (
	"minishop/pkg/infrastructure/response"

	"github.com/gofiber/fiber/v2"
)

type handlerImpl struct {
	serv ProductService
}

func (h *handlerImpl) GetAllProducts(c *fiber.Ctx) error {
	result, err := h.serv.GetAllProducts()
	if err != nil {
		return response.ExceptionError(c, err)
	}
	return response.SuccessResponse(c, 200, "get all products success", result)
}
func (h *handlerImpl) GetProductById(c *fiber.Ctx) error {
	productId := c.Params("productId")
	result, err := h.serv.GetProductById(&productId)
	if err != nil {
		return response.ExceptionError(c, err)
	}
	return response.SuccessResponse(c, 200, "get product success", result)
}
func (h *handlerImpl) InsertProduct(c *fiber.Ctx) error {
	body := new(ProductAddRequestDto)
	if err := c.BodyParser(&body); err != nil {
		return response.ErrorResponse(c, 400, "cannot parse to json", err.Error())
	}
	if err := h.serv.InsertProduct(body); err != nil {
		return response.ExceptionError(c, err)
	}
	return response.SuccessResponse(c, 201, "add product success", nil)
}
func (h *handlerImpl) UpdateProduct(c *fiber.Ctx) error {
	productId := c.Params("productId")
	body := new(ProductUpdateRequestDto)
	if err := c.BodyParser(&body); err != nil {
		return response.ErrorResponse(c, 400, "cannot parse to json", err.Error())
	}
	if err := h.serv.UpdateProduct(&productId, body); err != nil {
		return response.ExceptionError(c, err)
	}
	return response.SuccessResponse(c, 200, "update product success", nil)
}
func (h *handlerImpl) DeleteProduct(c *fiber.Ctx) error {
	productId := c.Params("productId")
	if err := h.serv.DeleteProductById(&productId); err != nil {
		return response.ExceptionError(c, err)
	}
	return response.SuccessResponse(c, 200, "delete product success", nil)
}
