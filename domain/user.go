package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	Name      string         `gorm:"not null" json:"name"`
	Address   string         `gorm:"not null" json:"address"`
	Phone     string         `gorm:"not null" json:"phone"`
	IdAccount uint           `gorm:"not null" json:"id_account"`
	Account   *Account       `json:"account" gorm:"foreignKey:IdAccount;references:ID"`
	CreatedAt *time.Time     `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type UserRepository interface {
	RetrieveAllUser() ([]User, error)
	RetrieveUserByID(id uint) (*User, error)
	CreateUser(User *User) (*User, error)
	UpdateUser(User *User) (*User, error)
	DeleteUser(id uint) error
}

type UserUseCase interface {
	FetchUsers(ctx context.Context) ([]User, error)
	FetchUserByID(ctx context.Context, id uint) (*User, error)
	CreateUser(ctx context.Context, req *User) (*User, error)
	UpdateUser(ctx context.Context, req *User) (*User, error)
	DeleteUser(ctx context.Context, id uint) error
	ShowUserLastNumber(ctx context.Context) (int, error)
}
