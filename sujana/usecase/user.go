package usecase

import (
	"context"
	"fmt"
	"sujana-be-web-go/domain"
	"time"
)

type userUseCase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewUserUseCase(user domain.UserRepository, t time.Duration) domain.UserUseCase {
	return &userUseCase{
		userRepository: user,
		contextTimeout: t,
	}
}

func (c *userUseCase) FetchUserByID(ctx context.Context, id uint) (*domain.User, error) {
	res, err := c.userRepository.RetrieveUserByID(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *userUseCase) FetchUsers(ctx context.Context) ([]domain.User, error) {
	res, err := c.userRepository.RetrieveAllUser()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *userUseCase) CreateUser(ctx context.Context, req *domain.User) (*domain.User, error) {
	res, err := c.userRepository.CreateUser(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *userUseCase) UpdateUser(ctx context.Context, req *domain.User) (*domain.User, error) {
	res, err := c.userRepository.UpdateUser(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *userUseCase) DeleteUser(ctx context.Context, id uint) error {
	err := c.userRepository.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}

func (c *userUseCase) ShowUserLastNumber(ctx context.Context) (int, error) {
	var res []domain.User
	res, err := c.userRepository.RetrieveAllUser()
	if err != nil {
		return 0, err
	}
	lastNumber := 0
	for _, v := range res {
		if v.ID > uint(lastNumber) {
			lastNumber = int(v.ID)
		}
	}
	fmt.Println(lastNumber)
	return lastNumber, nil
}
