package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/monforje/user-service/internal/entity"
	"github.com/monforje/user-service/internal/repository/postgres"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) (int64, error)
	GetByTelegramID(ctx context.Context, telegramID int64) (*entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User) error
	DeleteUser(ctx context.Context, telegramID int64) error
	ExistByTelegramID(ctx context.Context, telegramID int64) (bool, error)
}

type Repository struct {
	User UserRepository
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		User: postgres.NewUserRepo(db),
	}
}
