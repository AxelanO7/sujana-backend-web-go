package repository

import (
	"errors"
	"fmt"
	"sujana-be-web-go/domain"

	"gorm.io/gorm"
)

type posgreUserRepository struct {
	DB *gorm.DB
}

func NewPostgreUser(client *gorm.DB) domain.UserRepository {
	return &posgreUserRepository{
		DB: client,
	}
}

func (a *posgreUserRepository) RetrieveAllUser() ([]domain.User, error) {
	var res []domain.User
	err := a.DB.
		Model(domain.User{}).
		Preload("User").
		Find(&res).Error
	if err != nil {
		return []domain.User{}, err
	}
	fmt.Println(res)
	return res, nil
}

func (a *posgreUserRepository) RetrieveUserByID(id uint) (*domain.User, error) {
	var res domain.User
	err := a.DB.
		Model(domain.User{}).
		Where("id = ?", id).
		Preload("User").
		Take(&res).Error
	if err != nil {
		return &domain.User{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &domain.User{}, fmt.Errorf("record not found")
	}
	fmt.Println(res)
	return &res, nil
}

func (a *posgreUserRepository) CreateUser(user *domain.User) (*domain.User, error) {
	err := a.DB.
		Model(domain.User{}).
		Create(user).Error
	if err != nil {
		return &domain.User{}, err
	}
	fmt.Println(user)
	return user, nil
}

func (a *posgreUserRepository) UpdateUser(user *domain.User) (*domain.User, error) {
	err := a.DB.
		Model(domain.User{}).
		Where("id = ?", user.ID).
		Updates(user).Error
	if err != nil {
		return &domain.User{}, err
	}
	fmt.Println(user)
	return user, nil
}

func (a *posgreUserRepository) DeleteUser(id uint) error {
	err := a.DB.
		Model(domain.User{}).
		Where("id = ?", id).
		Delete(&domain.User{}).Error
	if err != nil {
		return err
	}
	return nil
}
