package domain

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type Account struct {
	ID        uint           `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	Username  *string        `json:"username"`
	Password  string         `gorm:"not null" json:"password"`
	Role      string         `gorm:"not null" json:"role"`
	CreatedAt *time.Time     `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type TokenClaims struct {
	Account *Account `json:"account"`
	jwt.StandardClaims
}

type LoginPayload struct {
	Username *string `json:"username"`
	Password string  `json:"password"`
}

type AccountRepository interface {
	RetrieveAllAccount() ([]Account, error)
	RetrieveAccountByID(id uint) (*Account, error)
	CreateAccount(account *Account) (*Account, error)
	UpdateAccount(account *Account) (*Account, error)
	DeleteAccount(id uint) error
	RetrieveByUsername(username string) (*Account, error)
}

type AccountUseCase interface {
	FetchAccounts(ctx context.Context) ([]Account, error)
	FetchAccountByID(ctx context.Context, id uint) (*Account, error)
	CreateAccount(ctx context.Context, req *Account) (*Account, error)
	UpdateAccount(ctx context.Context, req *Account) (*Account, error)
	DeleteAccount(ctx context.Context, id uint) error
	LoginAccount(ctx context.Context, req *LoginPayload) (*Account, string, error)
	RegisterAccount(ctx context.Context, req *LoginPayload) (*Account, error)
}
