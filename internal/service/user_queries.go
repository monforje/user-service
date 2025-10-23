package service

import (
	"context"
	"log"

	"github.com/monforje/user-service/internal/entity"
	validator "github.com/monforje/user-service/pkg/validator"
)

func (s *userService) GetByTelegramID(ctx context.Context, telegramID int64) (*entity.User, error) {
	if err := validator.ValidateTelegramID(telegramID); err != nil {
		log.Println(err)
		return nil, err
	}

	user, err := s.repository.GetByTelegramID(ctx, telegramID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) IsExist(ctx context.Context, telegramID int64) (bool, error) {
	if err := validator.ValidateTelegramID(telegramID); err != nil {
		log.Println(err)
		return false, err
	}

	return s.repository.ExistByTelegramID(ctx, telegramID)
}
