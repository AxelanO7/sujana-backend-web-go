package repository

import (
	"fmt"
	"sujana-be-web-go/domain"

	"gorm.io/gorm"
)

type posgreOrderRepository struct {
	DB *gorm.DB
}

func NewPostgreOrder(client *gorm.DB) domain.OrderRepository {
	return &posgreOrderRepository{
		DB: client,
	}
}

func (a *posgreOrderRepository) RetrieveOrders() ([]domain.Order, error) {
	var res []domain.Order
	err := a.DB.
		Model(domain.Order{}).
		Find(&res).Error
	if err != nil {
		return []domain.Order{}, err
	}
	fmt.Println(res)
	return res, nil
}

func (a *posgreOrderRepository) CreateOrder(order domain.Order) (domain.Order, error) {
	err := a.DB.
		Model(domain.Order{}).
		Create(&order).Error
	if err != nil {
		return domain.Order{}, err
	}
	fmt.Println(order)
	return order, nil
}

func (a *posgreOrderRepository) RetrieveOrderById(id string) (domain.Order, error) {
	var res domain.Order
	err := a.DB.
		Model(domain.Order{}).
		Where("id = ?", id).
		First(&res).Error
	if err != nil {
		return domain.Order{}, err
	}
	fmt.Println(res)
	return res, nil
}

func (a *posgreOrderRepository) UpdateOrderById(in domain.Order) (domain.Order, error) {
	err := a.DB.
		Model(domain.Order{}).
		Where("id = ?", in.ID).
		Updates(&in).Error
	if err != nil {
		return domain.Order{}, err
	}
	fmt.Println(in)
	return in, nil
}

func (a *posgreOrderRepository) RemoveOrderById(id string) error {
	err := a.DB.
		Model(domain.Order{}).
		Where("id = ?", id).
		Delete(&domain.Order{}).Error
	if err != nil {
		return err
	}
	return nil
}
