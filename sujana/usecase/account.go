package usecase

import (
	"context"
	"fmt"
	"sujana-be-web-go/domain"
	"sujana-be-web-go/middleware"
	"sujana-be-web-go/utils"
	"time"
)

type accountUseCase struct {
	accountRepository domain.AccountRepository
	contextTimeout    time.Duration
}

func NewAccountUseCase(account domain.AccountRepository, t time.Duration) domain.AccountUseCase {
	return &accountUseCase{
		accountRepository: account,
		contextTimeout:    t,
	}
}

func (c *accountUseCase) RegisterAccount(ctx context.Context, req *domain.LoginPayload) (*domain.Account, error) {
	password, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("unable to hash password: %v", err)
	}
	req.Password = password
	res, err := c.accountRepository.CreateAccount(&domain.Account{
		Username: req.Username,
		Password: req.Password,
		Role:     "user",
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *accountUseCase) LoginAccount(ctx context.Context, req *domain.LoginPayload) (*domain.Account, string, error) {
	res, err := c.accountRepository.RetrieveByUsername(*req.Username)
	if err != nil {
		return nil, "", err
	}
	err = utils.VerifyPassword(req.Password, res.Password)
	if err != nil {
		return nil, "", fmt.Errorf("error verifying password: %v", err)
	}
	tokPay := domain.TokenClaims{
		Account: res,
	}
	token, err := middleware.CreateToken(&tokPay)
	if err != nil {
		return nil, "", fmt.Errorf("cannot create token: %v", err)
	}
	return res, token, nil
}

func (c *accountUseCase) FetchAccountByID(ctx context.Context, id uint) (*domain.Account, error) {
	res, err := c.accountRepository.RetrieveAccountByID(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *accountUseCase) FetchAccounts(ctx context.Context) ([]domain.Account, error) {
	res, err := c.accountRepository.RetrieveAllAccount()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *accountUseCase) CreateAccount(ctx context.Context, req *domain.Account) (*domain.Account, error) {
	password, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("unable to hash password: %v", err)
	}
	req.Password = password
	res, err := c.accountRepository.CreateAccount(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *accountUseCase) UpdateAccount(ctx context.Context, req *domain.Account) (*domain.Account, error) {
	res, err := c.accountRepository.UpdateAccount(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *accountUseCase) DeleteAccount(ctx context.Context, id uint) error {
	err := c.accountRepository.DeleteAccount(id)
	if err != nil {
		return err
	}
	return nil
}
