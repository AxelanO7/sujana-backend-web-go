package usecase

import (
	"context"
	"sujana-be-web-go/domain"
	"time"
)

type orderUseCase struct {
	orderRepository domain.OrderRepository
	contextTimeout  time.Duration
}

func NewOrderUseCase(order domain.OrderRepository, t time.Duration) domain.OrderUseCase {
	return &orderUseCase{
		orderRepository: order,
		contextTimeout:  t,
	}
}

func (c *orderUseCase) ShowOrders(ctx context.Context) ([]domain.Order, error) {
	return c.orderRepository.RetrieveOrders()
}

func (c *orderUseCase) AddOrder(ctx context.Context, order domain.Order) (domain.Order, error) {
	return c.orderRepository.CreateOrder(order)
}

func (c *orderUseCase) ShowOrderById(ctx context.Context, id string) (domain.Order, error) {
	return c.orderRepository.RetrieveOrderById(id)
}

func (c *orderUseCase) EditOrderById(ctx context.Context, order domain.Order) (domain.Order, error) {
	return c.orderRepository.UpdateOrderById(order)
}

func (c *orderUseCase) DeleteOrderById(ctx context.Context, id string) error {
	return c.orderRepository.RemoveOrderById(id)
}
