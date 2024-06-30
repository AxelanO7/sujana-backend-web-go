package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID          uint           `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	OrderID     uint           `json:"order_id"`
	Name        string         `json:"name"`
	Type        string         `json:"type"`
	TotalPeople int16          `json:"total_people"`
	TotalDay    int16          `json:"total_day"`
	Price       float64        `json:"price"`
	TotalPrice  float64        `json:"total_price"`
	Phone       string         `json:"phone"`
	CreatedAt   *time.Time     `json:"created_at"`
	UpdatedAt   *time.Time     `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type OrderRepository interface {
	RetrieveOrders() ([]Order, error)
	RetrieveOrderById(id string) (Order, error)
	CreateOrder(in Order) (Order, error)
	UpdateOrderById(in Order) (Order, error)
	RemoveOrderById(id string) error
}

type OrderUseCase interface {
	ShowOrders(ctx context.Context) ([]Order, error)
	ShowOrderById(ctx context.Context, id string) (Order, error)
	AddOrder(ctx context.Context, in Order) (Order, error)
	EditOrderById(ctx context.Context, in Order) (Order, error)
	DeleteOrderById(ctx context.Context, id string) error
}
