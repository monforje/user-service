package service

import (
	"context"
	"log"

	"github.com/monforje/user-service/internal/entity"
	validator "github.com/monforje/user-service/pkg/validator"
)

func (s *userService) Create(ctx context.Context, user *entity.User) (int64, error) {
	if err := validator.ValidateTelegramID(user.TelegramID); err != nil {
		log.Println(err)
		return 0, err
	}

	if err := validator.ValidatePhone(user.Phone); err != nil {
		log.Println(err)
		return 0, err
	}

	if err := validator.ValidateUsername(user.Username); err != nil {
		log.Println(err)
		return 0, err
	}

	id, err := s.repository.CreateUser(ctx, user)
	if err != nil {
		return 0, err
	}

	return id, nil

}

func (s *userService) Update(ctx context.Context, user *entity.User) error {
	if err := validator.ValidateTelegramID(user.TelegramID); err != nil {
		log.Println(err)
		return err
	}

	if err := validator.ValidatePhone(user.Phone); err != nil {
		log.Println(err)
		return err
	}

	if err := validator.ValidateUsername(user.Username); err != nil {
		log.Println(err)
		return err
	}

	if err := s.repository.UpdateUser(ctx, user); err != nil {
		return err
	}

	return nil
}

func (s *userService) Delete(ctx context.Context, telegramID int64) error {
	if err := validator.ValidateTelegramID(telegramID); err != nil {
		log.Println(err)
		return err
	}

	if err := s.repository.DeleteUser(ctx, telegramID); err != nil {
		return err
	}

	return nil
}
