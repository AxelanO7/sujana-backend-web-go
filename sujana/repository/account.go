package repository

import (
	"errors"
	"fmt"
	"sujana-be-web-go/domain"

	"gorm.io/gorm"
)

type posgreAccountRepository struct {
	DB *gorm.DB
}

func NewPostgreAccount(client *gorm.DB) domain.AccountRepository {
	return &posgreAccountRepository{
		DB: client,
	}
}

func (a *posgreAccountRepository) RetrieveAllAccount() ([]domain.Account, error) {
	var res []domain.Account
	err := a.DB.
		Model(domain.Account{}).
		Find(&res).Error
	if err != nil {
		return []domain.Account{}, err
	}
	fmt.Println(res)
	return res, nil
}

func (a *posgreAccountRepository) RetrieveByUsername(accountname string) (*domain.Account, error) {
	var res domain.Account
	err := a.DB.
		Model(domain.Account{}).
		Where("username = ?", accountname).
		Take(&res).Error
	if err != nil {
		return &domain.Account{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &domain.Account{}, fmt.Errorf("record not found")
	}
	fmt.Println(res)
	return &res, nil
}

func (a *posgreAccountRepository) RetrieveAccountByID(id uint) (*domain.Account, error) {
	var res domain.Account
	err := a.DB.
		Model(domain.Account{}).
		Where("id = ?", id).
		Take(&res).Error
	if err != nil {
		return &domain.Account{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &domain.Account{}, fmt.Errorf("record not found")
	}
	fmt.Println(res)
	return &res, nil
}

func (a *posgreAccountRepository) CreateAccount(account *domain.Account) (*domain.Account, error) {
	err := a.DB.
		Model(domain.Account{}).
		Create(account).Error
	if err != nil {
		return &domain.Account{}, err
	}
	fmt.Println(account)
	return account, nil
}

func (a *posgreAccountRepository) UpdateAccount(account *domain.Account) (*domain.Account, error) {
	err := a.DB.
		Model(domain.Account{}).
		Where("id = ?", account.ID).
		Updates(account).Error
	if err != nil {
		return &domain.Account{}, err
	}
	fmt.Println(account)
	return account, nil
}

func (a *posgreAccountRepository) DeleteAccount(id uint) error {
	err := a.DB.
		Model(domain.Account{}).
		Where("id = ?", id).
		Delete(&domain.Account{}).Error
	if err != nil {
		return err
	}
	return nil
}
