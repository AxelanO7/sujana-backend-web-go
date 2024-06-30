package delivery

import (
	"sujana-be-web-go/domain"

	"github.com/gofiber/fiber/v2"
)

type OrderHandler struct {
	OrderUC domain.OrderUseCase
}

func NewOrderHandler(c *fiber.App, das domain.OrderUseCase) {
	handler := &OrderHandler{
		OrderUC: das,
	}
	api := c.Group("/order")

	public := api.Group("/public")
	public.Get("/package/", handler.ShowOrders)
	public.Post("/package/", handler.AddOrder)
	public.Get("/package/:id", handler.ShowOrderById)
	public.Put("/package/:id", handler.EditOrderById)
	public.Delete("/package/:id", handler.DeleteOrderById)

	_ = api.Group("/private")
}

func (t *OrderHandler) ShowOrders(c *fiber.Ctx) error {
	res, er := t.OrderUC.ShowOrders(c.Context())
	if er != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"success": false,
			"data":    nil,
			"message": er.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"success": true,
		"data":    res,
		"message": "Success get data",
	})
}

func (t *OrderHandler) ShowOrderById(c *fiber.Ctx) error {
	id := c.Params("id")
	res, er := t.OrderUC.ShowOrderById(c.Context(), id)
	if er != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"success": false,
			"data":    nil,
			"message": er.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"success": true,
		"data":    res,
		"message": "Success get data",
	})
}

func (t *OrderHandler) AddOrder(c *fiber.Ctx) error {
	var in domain.Order
	if err := c.BodyParser(&in); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"success": false,
			"data":    nil,
			"message": err.Error(),
		})
	}

	res, er := t.OrderUC.AddOrder(c.Context(), in)
	if er != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"success": false,
			"data":    nil,
			"message": er.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  201,
		"success": true,
		"data":    res,
		"message": "Success create data",
	})
}

func (t *OrderHandler) EditOrderById(c *fiber.Ctx) error {
	var in domain.Order
	if err := c.BodyParser(&in); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"success": false,
			"data":    nil,
			"message": err.Error(),
		})
	}

	res, er := t.OrderUC.EditOrderById(c.Context(), in)
	if er != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"success": false,
			"data":    nil,
			"message": er.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"success": true,
		"data":    res,
		"message": "Success update data",
	})
}

func (t *OrderHandler) DeleteOrderById(c *fiber.Ctx) error {
	id := c.Params("id")
	er := t.OrderUC.DeleteOrderById(c.Context(), id)
	if er != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"success": false,
			"data":    nil,
			"message": er.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"success": true,
		"data":    nil,
		"message": "Success delete data",
	})
}
