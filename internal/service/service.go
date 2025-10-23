package service

import (
	"context"

	"github.com/monforje/user-service/internal/entity"
	"github.com/monforje/user-service/internal/repository"
)

type UserService interface {
	Create(ctx context.Context, user *entity.User) (int64, error)
	GetByTelegramID(ctx context.Context, telegramID int64) (*entity.User, error)
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, telegramID int64) error
	IsExist(ctx context.Context, telegramID int64) (bool, error)
}

type Service struct {
	User UserService
}

func New(repository *repository.Repository) *Service {
	return &Service{
		User: newUserService(repository.User),
	}
}
